package namespaceUtils

import (
	"fmt"
	"strings"

	"go.temporal.io/api/enums/v1"
)

// CreateIndexedValueTypeMap is a helper function takes in a map[string]string and returns another
// map where each entry value from the old map is translated from a plain string into an enums.IndexedValueType.
func CreateIndexedValueTypeMap(m *map[string]string) (map[string]enums.IndexedValueType, error) {
	specCustomSearchAttributes := make(map[string]enums.IndexedValueType, len(*m))
	for searchAttributeNameString, searchAttributeTypeString := range *m {
		indexedValueType, err := searchAttributeTypeStringToEnum(searchAttributeTypeString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse search attribute %s because its type is %s: %w", searchAttributeNameString, searchAttributeTypeString, err)
		}
		specCustomSearchAttributes[searchAttributeNameString] = indexedValueType
	}
	return specCustomSearchAttributes, nil
}

// searchAttributeTypeStringToEnum retrieves the actual IndexedValueType for a given string.
// It expects searchAttributeTypeString to be a string representation of the valid Go type.
// Returns the IndexedValueType if parsing is successful, otherwise an error.
// See https://docs.temporal.io/visibility#supported-types for supported types.
func searchAttributeTypeStringToEnum(searchAttributeTypeString string) (enums.IndexedValueType, error) {
	for k, v := range enums.IndexedValueType_shorthandValue {
		if strings.EqualFold(searchAttributeTypeString, k) {
			return enums.IndexedValueType(v), nil
		}
	}
	return enums.INDEXED_VALUE_TYPE_UNSPECIFIED, fmt.Errorf("unsupported search attribute type: %v", searchAttributeTypeString)
}
