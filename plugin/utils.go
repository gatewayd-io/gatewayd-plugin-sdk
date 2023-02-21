package plugin

import "google.golang.org/protobuf/types/known/structpb"

func GetAttr(req *structpb.Struct, key string, defaultValue interface{}) interface{} {
	if req == nil {
		return defaultValue
	}

	if value, ok := req.Fields[key]; ok {
		return value.AsInterface()
	}

	return defaultValue
}
