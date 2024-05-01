package normalizer

import (
	"fmt"
	"golang.org/x/exp/maps"
	"strconv"
	"strings"
)

func Normalize(header map[string]interface{}, body, payments, extensions []map[string]interface{}) string {
	packetWrapper := map[string]interface{}{
		"payments":   payments,
		"extensions": extensions,
	}
	for key, val := range header {
		if val != "" {
			if key == "Authorization" {
				val = val.(string)[7:]
			}
			packetWrapper[key] = val
		}
	}
	if len(body) > 1 {
		packetWrapper["packets"] = body
	} else {
		for k, v := range body[0] {
			packetWrapper[k] = v
		}
	}
	result := make(map[string]interface{})
	flatMap("", packetWrapper, result)
	return finalize(result)
}

func finalize(result map[string]interface{}) string {
	resultKeys := maps.Keys(result)
	var valueConv string
	var normalized []string
	for _, key := range resultKeys {
		if result[key] != nil && result[key] != "" {
			switch result[key].(type) {
			case int:
				valueConv = strconv.Itoa(result[key].(int))
			case float64:
				valueConv = strconv.FormatFloat(result[key].(float64), 'f', -1, 64)
			default:
				valueConv = result[key].(string)
			}
			strings.Replace(valueConv, "#", "##", -1)
			normalized = append(normalized, valueConv)
		} else {
			normalized = append(normalized, "#")
		}
	}
	return strings.Join(normalized, "#")
}

func flatMap(rootKey string, data interface{}, result map[string]interface{}) {
	if data == nil {
		result[rootKey] = ""
	}
	switch data.(type) {
	case nil:
		result[rootKey] = ""
	case map[string]interface{}:
		for k, v := range data.(map[string]interface{}) {
			k = generateKey(rootKey, k)
			flatMap(k, v, result)
		}
	case []map[string]interface{}:
		for index, i := range data.([]map[string]interface{}) {
			flatMap(fmt.Sprintf("E%d", index), i, result)
		}
	case int:
		result[rootKey] = strconv.Itoa(data.(int))
	case int64:
		result[rootKey] = strconv.Itoa(int(data.(int64)))
	case float64:
		result[rootKey] = strconv.FormatFloat(data.(float64), 'f', -1, 64)
	case []interface{}:
		for index, i := range data.([]interface{}) {
			flatMap(fmt.Sprintf("E%d", index), i, result)
		}
	case bool:
		result[rootKey] = strconv.FormatBool(data.(bool))
	default:
		if data.(string) == "" {
			result[rootKey] = ""
		} else {
			result[rootKey] = data.(string)
		}
	}
}

func generateKey(root string, key string) string {
	if root != "" {
		return fmt.Sprintf("%s.%s", root, key)
	} else {
		return key
	}
}
