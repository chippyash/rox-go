package extensions_test

import (
	"github.com/rollout/rox-go/v4/core/context"
	"github.com/rollout/rox-go/v4/core/extensions"
	"github.com/rollout/rox-go/v4/core/properties"
	"github.com/rollout/rox-go/v4/core/repositories"
	"github.com/rollout/rox-go/v4/core/roxx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPropertiesExtensionsRoxxPropertiesExtensionsString(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewStringProperty("testKey", "test"))

	assert.Equal(t, true, parser.EvaluateExpression(`eq("test", property("testKey"))`, nil).Value())
}

func TestPropertiesExtensionsRoxxPropertiesExtensionsInt(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewIntegerProperty("testKey", 3))

	assert.Equal(t, true, parser.EvaluateExpression(`eq(3, property("testKey"))`, nil).Value())
}

func TestPropertiesExtensionsRoxxPropertiesExtensionsFloat(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewFloatProperty("testKey", 3.3))

	assert.Equal(t, true, parser.EvaluateExpression(`eq(3.3, property("testKey"))`, nil).Value())
}

func TestPropertiesExtensionsRoxxPropertiesExtensionsWithContextString(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewComputedStringProperty("CustomPropertyTestKey", func(context context.Context) string {
		return context.Get("ContextTestKey").(string)
	}))

	ctx := context.NewContext(map[string]interface{}{"ContextTestKey": "test"})
	assert.Equal(t, true, parser.EvaluateExpression(`eq("test", property("CustomPropertyTestKey"))`, ctx).Value())
}

func TestPropertiesExtensionsRoxxPropertiesExtensionsWithContextInt(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewComputedIntegerProperty("CustomPropertyTestKey", func(context context.Context) int {
		return context.Get("ContextTestKey").(int)
	}))

	ctx := context.NewContext(map[string]interface{}{"ContextTestKey": 3})
	assert.Equal(t, true, parser.EvaluateExpression(`eq(3, property("CustomPropertyTestKey"))`, ctx).Value())
}

func TestPropertiesExtensionsRoxxPropertiesExtensionsWithContextIntWithString(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewComputedIntegerProperty("CustomPropertyTestKey", func(context context.Context) int {
		return context.Get("ContextTestKey").(int)
	}))

	ctx := context.NewContext(map[string]interface{}{"ContextTestKey": 3})
	assert.Equal(t, false, parser.EvaluateExpression(`eq("3", property("CustomPropertyTestKey"))`, ctx).Value())
}

func TestPropertiesExtensionsRoxxPropertiesExtensionsWithContextIntNotEqual(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewComputedIntegerProperty("CustomPropertyTestKey", func(context context.Context) int {
		return context.Get("ContextTestKey").(int)
	}))

	ctx := context.NewContext(map[string]interface{}{"ContextTestKey": 3})
	assert.Equal(t, false, parser.EvaluateExpression(`eq(4, property("CustomPropertyTestKey"))`, ctx).Value())
}

func TestPropertiesExtensionsUnknownProperty(t *testing.T) {
	customPropertiesRepository := repositories.NewCustomPropertyRepository()
	parser := roxx.NewParser()
	extensions.NewPropertiesExtensions(parser, customPropertiesRepository).Extend()

	customPropertiesRepository.AddCustomProperty(properties.NewStringProperty("testKey", "test"))

	assert.Equal(t, false, parser.EvaluateExpression(`eq("test", property("testKey1"))`, nil).Value())
}
