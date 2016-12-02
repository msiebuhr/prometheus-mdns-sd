// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// (Morten Siebuhr: This code is taken from https://github.com/prometheus/prometheus/config/config.go)

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/prometheus/common/model"
	//"gopkg.in/yaml.v2"
)

func checkOverflow(m map[string]interface{}, ctx string) error {
	if len(m) > 0 {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		return fmt.Errorf("unknown fields in %s: %s", ctx, strings.Join(keys, ", "))
	}
	return nil
}

// TargetGroup is a set of targets with a common label set.
type TargetGroup struct {
	// Targets is a list of targets identified by a label set. Each target is
	// uniquely identifiable in the group by its address label.
	Targets []model.LabelSet `json:"targets"`
	// Labels is a set of labels that is common across all targets in the group.
	Labels model.LabelSet `json:"labels"`

	// Source is an identifier that describes a group of targets.
	Source string
}

func (tg TargetGroup) String() string {
	return tg.Source
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (tg *TargetGroup) UnmarshalYAML(unmarshal func(interface{}) error) error {
	g := struct {
		Targets []string               `yaml:"targets"`
		Labels  model.LabelSet         `yaml:"labels"`
		XXX     map[string]interface{} `yaml:",inline"`
	}{}
	if err := unmarshal(&g); err != nil {
		return err
	}
	tg.Targets = make([]model.LabelSet, 0, len(g.Targets))
	for _, t := range g.Targets {
		tg.Targets = append(tg.Targets, model.LabelSet{
			model.AddressLabel: model.LabelValue(t),
		})
	}
	tg.Labels = g.Labels
	return checkOverflow(g.XXX, "target_group")
}

// MarshalYAML implements the yaml.Marshaler interface.
func (tg TargetGroup) MarshalYAML() (interface{}, error) {
	g := &struct {
		Targets []string       `yaml:"targets"`
		Labels  model.LabelSet `yaml:"labels,omitempty"`
	}{
		Targets: make([]string, 0, len(tg.Targets)),
		Labels:  tg.Labels,
	}
	for _, t := range tg.Targets {
		g.Targets = append(g.Targets, string(t[model.AddressLabel]))
	}
	return g, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tg *TargetGroup) UnmarshalJSON(b []byte) error {
	g := struct {
		Targets []string       `json:"targets"`
		Labels  model.LabelSet `json:"labels"`
	}{}
	if err := json.Unmarshal(b, &g); err != nil {
		return err
	}
	tg.Targets = make([]model.LabelSet, 0, len(g.Targets))
	for _, t := range g.Targets {
		tg.Targets = append(tg.Targets, model.LabelSet{
			model.AddressLabel: model.LabelValue(t),
		})
	}
	tg.Labels = g.Labels
	return nil
}
