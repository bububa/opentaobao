package aesolution

// SingleItemResponseDto 结构体
type SingleItemResponseDto struct {
	// Execution result of each item
	ItemExecutionResult string `json:"item_execution_result,omitempty" xml:"item_execution_result,omitempty"`
	// Corresponding to the item_content_id defined by the seller when invoking the API: aliexpress.solution.feed.submit
	ItemContentId string `json:"item_content_id,omitempty" xml:"item_content_id,omitempty"`
}
