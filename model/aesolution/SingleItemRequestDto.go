package aesolution

// SingleItemRequestDTO 
type SingleItemRequestDTO struct {
    // Content of each item, which follows different format according to different feed type.
    ItemContent   string `json:"item_content,omitempty" xml:"item_content,omitempty"`
    // The id of the item_content, which could be defined by the seller. item_content_id should be unique among all the items in item_list.This field also appears in the API:aliexpress.solution.feed.query, which is regarding the convenience for the sellers to match the item_execuation_result with the item_content.
    ItemContentId   string `json:"item_content_id,omitempty" xml:"item_content_id,omitempty"`
}
