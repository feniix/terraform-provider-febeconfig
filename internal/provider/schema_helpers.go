package provider

// stolen from https://github.com/hashicorp/terraform-provider-kubernetes/blob/master/kubernetes/schema_helpers.go

func conditionalDefault(condition bool, defaultValue interface{}) interface{} {
	if !condition {
		return nil
	}

	return defaultValue
}

