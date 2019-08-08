package config

import (
	"github.com/prometheus/prometheus/pkg/rulefmt"
)

func LoadRules(paths []string) ([]rulefmt.RuleGroups, []error) {
	rules := []rulefmt.RuleGroups{}
	for _, p := range paths {
		ruleGroups, _ := rulefmt.ParseFile(p)
		rules = append(rules, *ruleGroups)
	}

	return rules, nil

}
