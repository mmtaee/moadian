package normalizer

func DTO(data map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	if header, ok := data["header"]; ok {
		result["header"] = mapDto(header.(map[string]interface{}))
	}
	result["body"] = sliceDto(data["body"])
	result["payments"] = sliceDto(data["payments"])
	result["extensions"] = sliceDto(data["extensions"])
	return result
}

func mapDto(m map[string]interface{}) map[string]interface{} {
	dto := make(map[string]interface{})
	for key, val := range m {
		if val != nil || val != "" {
			dto[key] = val
		}
	}
	return dto
}

func sliceDto(s interface{}) []map[string]interface{} {
	if s == nil {
		return nil
	}
	original := s.([]map[string]interface{})
	dto := make([]map[string]interface{}, len(original))
	for i, m := range original {
		dto[i] = mapDto(m)
	}
	return dto
}
