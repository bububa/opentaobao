package tmallservice

// WorkerDto 
type WorkerDto struct {

    // 工人居住地址
    Address   *AddressDto `json:"address,omitempty"`

    // 111
    IdentityId   string `json:"identity_id,omitempty"`

    // 11
    Name   string `json:"name,omitempty"`

    // 1
    Phone   int64 `json:"phone,omitempty"`

    // 111
    ProviderId   int64 `json:"provider_id,omitempty"`

    // 1111
    ProviderName   string `json:"provider_name,omitempty"`

    // 111
    RegisterTime   string `json:"register_time,omitempty"`

    // 1111
    ServiceAreas   []DivisionDto `json:"service_areas,omitempty"`

    // 11
    ServiceTypes   []String `json:"service_types,omitempty"`

    // 1111
    WorkType   string `json:"work_type,omitempty"`

    // 111
    Photo   string `json:"photo,omitempty"`

    // 11
    IdCardPicBack   string `json:"id_card_pic_back,omitempty"`

    // 11
    IdCardPic   string `json:"id_card_pic,omitempty"`

    // 111
    HandheldCardPic   string `json:"handheld_card_pic,omitempty"`

    // 工人所属行业类型
    BizType   string `json:"biz_type,omitempty"`

    // 覆盖的service_code列表，|隔开
    ServiceCodes   string `json:"service_codes,omitempty"`

    // 工人支持的商品类目，格式：类目id1|类目id2
    CoverCategoryIds   string `json:"cover_category_ids,omitempty"`

    // 网点编码
    ServiceStoreCode   string `json:"service_store_code,omitempty"`

}
