// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeSlice_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_endpoint", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("endpoint"); err == nil {
				assert.Equal(t, string(defaultConfig.Endpoint.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Endpoint.String()

			cmdFlags.Set("endpoint", testValue)
			if vString, err := cmdFlags.GetString("endpoint"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Endpoint)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_commandApiPath", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("commandApiPath"); err == nil {
				assert.Equal(t, string(defaultConfig.CommandAPIPath.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.CommandAPIPath.String()

			cmdFlags.Set("commandApiPath", testValue)
			if vString, err := cmdFlags.GetString("commandApiPath"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.CommandAPIPath)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_analyzeLinkPath", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("analyzeLinkPath"); err == nil {
				assert.Equal(t, string(defaultConfig.AnalyzeLinkPath.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.AnalyzeLinkPath.String()

			cmdFlags.Set("analyzeLinkPath", testValue)
			if vString, err := cmdFlags.GetString("analyzeLinkPath"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AnalyzeLinkPath)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_quboleTokenKey", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("quboleTokenKey"); err == nil {
				assert.Equal(t, string(defaultConfig.TokenKey), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("quboleTokenKey", testValue)
			if vString, err := cmdFlags.GetString("quboleTokenKey"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.TokenKey)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_defaultClusterLabel", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("defaultClusterLabel"); err == nil {
				assert.Equal(t, string(defaultConfig.DefaultClusterLabel), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("defaultClusterLabel", testValue)
			if vString, err := cmdFlags.GetString("defaultClusterLabel"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.DefaultClusterLabel)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_readRateLimiter.qps", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("readRateLimiter.qps"); err == nil {
				assert.Equal(t, int(defaultConfig.ReadRateLimiter.QPS), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("readRateLimiter.qps", testValue)
			if vInt, err := cmdFlags.GetInt("readRateLimiter.qps"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.ReadRateLimiter.QPS)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_readRateLimiter.burst", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("readRateLimiter.burst"); err == nil {
				assert.Equal(t, int(defaultConfig.ReadRateLimiter.Burst), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("readRateLimiter.burst", testValue)
			if vInt, err := cmdFlags.GetInt("readRateLimiter.burst"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.ReadRateLimiter.Burst)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_writeRateLimiter.qps", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("writeRateLimiter.qps"); err == nil {
				assert.Equal(t, int(defaultConfig.WriteRateLimiter.QPS), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("writeRateLimiter.qps", testValue)
			if vInt, err := cmdFlags.GetInt("writeRateLimiter.qps"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.WriteRateLimiter.QPS)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_writeRateLimiter.burst", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("writeRateLimiter.burst"); err == nil {
				assert.Equal(t, int(defaultConfig.WriteRateLimiter.Burst), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("writeRateLimiter.burst", testValue)
			if vInt, err := cmdFlags.GetInt("writeRateLimiter.burst"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.WriteRateLimiter.Burst)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_caching.size", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("caching.size"); err == nil {
				assert.Equal(t, int(defaultConfig.Caching.Size), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("caching.size", testValue)
			if vInt, err := cmdFlags.GetInt("caching.size"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Caching.Size)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_caching.resyncInterval", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vString, err := cmdFlags.GetString("caching.resyncInterval"); err == nil {
				assert.Equal(t, string(defaultConfig.Caching.ResyncInterval.String()), vString)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := defaultConfig.Caching.ResyncInterval.String()

			cmdFlags.Set("caching.resyncInterval", testValue)
			if vString, err := cmdFlags.GetString("caching.resyncInterval"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.Caching.ResyncInterval)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_caching.workers", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			// Test that default value is set properly
			if vInt, err := cmdFlags.GetInt("caching.workers"); err == nil {
				assert.Equal(t, int(defaultConfig.Caching.Workers), vInt)
			} else {
				assert.FailNow(t, err.Error())
			}
		})

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("caching.workers", testValue)
			if vInt, err := cmdFlags.GetInt("caching.workers"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.Caching.Workers)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
