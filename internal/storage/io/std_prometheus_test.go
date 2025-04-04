package io_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/prometheus/prometheus/model/rulefmt"
	"github.com/stretchr/testify/assert"

	"github.com/slok/sloth/internal/log"
	"github.com/slok/sloth/internal/storage/io"
	"github.com/slok/sloth/pkg/common/model"
)

func TestGroupedRulesYAMLRepoStore(t *testing.T) {
	tests := map[string]struct {
		slos    []io.StdPrometheusStorageSLO
		expYAML string
		expErr  bool
	}{
		"Having 0 SLO rules should fail.": {
			slos:   []io.StdPrometheusStorageSLO{},
			expErr: true,
		},

		"Having 0 SLO rules generated should fail.": {
			slos: []io.StdPrometheusStorageSLO{
				{},
			},
			expErr: true,
		},

		"Having a single SLI recording rule should render correctly.": {
			slos: []io.StdPrometheusStorageSLO{
				{
					SLO: model.PromSLO{ID: "test1"},
					Rules: model.PromSLORules{
						SLIErrorRecRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Record: "test:record",
								Expr:   "test-expr",
								Labels: map[string]string{"test-label": "one"},
							},
						}},
					},
				},
			},
			expYAML: `
---
# Code generated by Sloth (dev): https://github.com/slok/sloth.
# DO NOT EDIT.

groups:
- name: sloth-slo-sli-recordings-test1
  rules:
  - record: test:record
    expr: test-expr
    labels:
      test-label: one
`,
		},
		"Having a single metadata recording rule should render correctly.": {
			slos: []io.StdPrometheusStorageSLO{
				{
					SLO: model.PromSLO{ID: "test1"},
					Rules: model.PromSLORules{
						MetadataRecRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Record: "test:record",
								Expr:   "test-expr",
								Labels: map[string]string{"test-label": "one"},
							},
						}},
					},
				},
			},
			expYAML: `
---
# Code generated by Sloth (dev): https://github.com/slok/sloth.
# DO NOT EDIT.

groups:
- name: sloth-slo-meta-recordings-test1
  rules:
  - record: test:record
    expr: test-expr
    labels:
      test-label: one
`,
		},
		"Having a single SLO alert rule should render correctly.": {
			slos: []io.StdPrometheusStorageSLO{
				{
					SLO: model.PromSLO{ID: "test1"},
					Rules: model.PromSLORules{
						AlertRules: model.PromRuleGroup{
							Interval: 42 * time.Minute,
							Rules: []rulefmt.Rule{
								{
									Alert:       "testAlert",
									Expr:        "test-expr",
									Labels:      map[string]string{"test-label": "one"},
									Annotations: map[string]string{"test-annot": "one"},
								},
							}},
					},
				},
			},
			expYAML: `
---
# Code generated by Sloth (dev): https://github.com/slok/sloth.
# DO NOT EDIT.

groups:
- name: sloth-slo-alerts-test1
  interval: 42m
  rules:
  - alert: testAlert
    expr: test-expr
    labels:
      test-label: one
    annotations:
      test-annot: one
`,
		},

		"Having a multiple SLO alert and recording rules should render correctly.": {
			slos: []io.StdPrometheusStorageSLO{
				{
					SLO: model.PromSLO{ID: "testa"},
					Rules: model.PromSLORules{
						SLIErrorRecRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Record: "test:record-a1",
								Expr:   "test-expr-a1",
								Labels: map[string]string{"test-label": "a-1"},
							},
							{
								Record: "test:record-a2",
								Expr:   "test-expr-a2",
								Labels: map[string]string{"test-label": "a-2"},
							},
						}},
						MetadataRecRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Record: "test:record-a3",
								Expr:   "test-expr-a3",
								Labels: map[string]string{"test-label": "a-3"},
							},
							{
								Record: "test:record-a4",
								Expr:   "test-expr-a4",
								Labels: map[string]string{"test-label": "a-4"},
							},
						}},
						AlertRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Alert:       "testAlertA1",
								Expr:        "test-expr-a1",
								Labels:      map[string]string{"test-label": "a-1"},
								Annotations: map[string]string{"test-annot": "a-1"},
							},
							{
								Alert:       "testAlertA2",
								Expr:        "test-expr-a2",
								Labels:      map[string]string{"test-label": "a-2"},
								Annotations: map[string]string{"test-annot": "a-2"},
							},
						}},
					},
				},
				{
					SLO: model.PromSLO{ID: "testb"},
					Rules: model.PromSLORules{
						SLIErrorRecRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Record: "test:record-b1",
								Expr:   "test-expr-b1",
								Labels: map[string]string{"test-label": "b-1"},
							},
						}},
						MetadataRecRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Record: "test:record-b2",
								Expr:   "test-expr-b2",
								Labels: map[string]string{"test-label": "b-2"},
							},
						}},
						AlertRules: model.PromRuleGroup{Rules: []rulefmt.Rule{
							{
								Alert:       "testAlertB1",
								Expr:        "test-expr-b1",
								Labels:      map[string]string{"test-label": "b-1"},
								Annotations: map[string]string{"test-annot": "b-1"},
							},
						}},
					},
				},
			},
			expYAML: `
---
# Code generated by Sloth (dev): https://github.com/slok/sloth.
# DO NOT EDIT.

groups:
- name: sloth-slo-sli-recordings-testa
  rules:
  - record: test:record-a1
    expr: test-expr-a1
    labels:
      test-label: a-1
  - record: test:record-a2
    expr: test-expr-a2
    labels:
      test-label: a-2
- name: sloth-slo-meta-recordings-testa
  rules:
  - record: test:record-a3
    expr: test-expr-a3
    labels:
      test-label: a-3
  - record: test:record-a4
    expr: test-expr-a4
    labels:
      test-label: a-4
- name: sloth-slo-alerts-testa
  rules:
  - alert: testAlertA1
    expr: test-expr-a1
    labels:
      test-label: a-1
    annotations:
      test-annot: a-1
  - alert: testAlertA2
    expr: test-expr-a2
    labels:
      test-label: a-2
    annotations:
      test-annot: a-2
- name: sloth-slo-sli-recordings-testb
  rules:
  - record: test:record-b1
    expr: test-expr-b1
    labels:
      test-label: b-1
- name: sloth-slo-meta-recordings-testb
  rules:
  - record: test:record-b2
    expr: test-expr-b2
    labels:
      test-label: b-2
- name: sloth-slo-alerts-testb
  rules:
  - alert: testAlertB1
    expr: test-expr-b1
    labels:
      test-label: b-1
    annotations:
      test-annot: b-1
`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			var gotYAML bytes.Buffer
			repo := io.NewStdPrometheusGroupedRulesYAMLRepo(&gotYAML, log.Noop)
			err := repo.StoreSLOs(context.TODO(), test.slos)

			if test.expErr {
				assert.Error(err)
			} else if assert.NoError(err) {
				assert.Equal(test.expYAML, gotYAML.String())
			}
		})
	}
}
