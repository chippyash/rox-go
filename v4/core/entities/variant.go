package entities

import (
	"github.com/rollout/rox-go/v4/core/context"
	"github.com/rollout/rox-go/v4/core/model"
	"github.com/rollout/rox-go/v4/core/roxx"
	"github.com/rollout/rox-go/v4/core/utils"
)

type internalVariant interface {
	Condition() string
	Parser() roxx.Parser
	ImpressionInvoker() model.ImpressionInvoker
	ClientExperiment() *model.Experiment
}

type variant struct {
	defaultValue      string
	options           []string
	condition         string
	parser            roxx.Parser
	globalContext     context.Context
	impressionInvoker model.ImpressionInvoker
	clientExperiment  *model.Experiment
	name              string
}

func NewVariant(defaultValue string, options []string) model.Variant {
	if options == nil {
		options = []string{}
	}
	allOptions := make([]string, len(options))
	copy(allOptions, options)
	if !utils.ContainsString(allOptions, defaultValue) {
		allOptions = append(allOptions, defaultValue)
	}

	return &variant{
		defaultValue: defaultValue,
		options:      allOptions,
	}
}

func (v *variant) DefaultValue() string {
	return v.defaultValue
}

func (v *variant) Options() []string {
	return v.options
}

func (v *variant) Name() string {
	return v.name
}

func (v *variant) SetForEvaluation(parser roxx.Parser, experiment *model.ExperimentModel, impressionInvoker model.ImpressionInvoker) {
	if experiment != nil {
		v.clientExperiment = model.NewExperiment(experiment)
		v.condition = experiment.Condition
	} else {
		v.clientExperiment = nil
		v.condition = ""
	}

	v.parser = parser
	v.impressionInvoker = impressionInvoker
}

func (v *variant) SetContext(globalContext context.Context) {
	v.globalContext = globalContext
}

func (v *variant) SetName(name string) {
	v.name = name
}

func (v *variant) GetValue(ctx context.Context) string {
	returnValue, _ := v.InternalGetValue(ctx)
	return returnValue
}

func (v *variant) InternalGetValue(ctx context.Context) (returnValue string, isDefault bool) {
	returnValue, isDefault = v.defaultValue, true
	mergedContext := context.NewMergedContext(v.globalContext, ctx)

	if v.parser != nil && v.condition != "" {
		evaluationResult := v.parser.EvaluateExpression(v.condition, mergedContext)
		value := evaluationResult.StringValue()
		if value != "" {
			returnValue, isDefault = value, false
		}
	}

	if v.impressionInvoker != nil {
		v.impressionInvoker.Invoke(model.NewReportingValue(v.name, returnValue), v.clientExperiment, mergedContext)
	}

	return returnValue, isDefault
}

func (v *variant) Condition() string {
	return v.condition
}

func (v *variant) Parser() roxx.Parser {
	return v.parser
}

func (v *variant) ImpressionInvoker() model.ImpressionInvoker {
	return v.impressionInvoker
}

func (v *variant) ClientExperiment() *model.Experiment {
	return v.clientExperiment
}
