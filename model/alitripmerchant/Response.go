package alitripmerchant

// Response 
type Response struct {
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 错误编码
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 返回结果
    Contents   []AddressSearchDto `json:"contents,omitempty" xml:"contents>address_search_dto,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 返回城市列表
    Content   *AddressListSearchDto `json:"content,omitempty" xml:"content,omitempty"`
    // 登出是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 是否注册成功
    RegisterResult   bool `json:"register_result,omitempty" xml:"register_result,omitempty"`
    // token
    Token   string `json:"token,omitempty" xml:"token,omitempty"`
    // offer列表
    Offers   []OfferDetailsDto `json:"offers,omitempty" xml:"offers>offer_details_dto,omitempty"`
    // 查询结果
    OrderDtos   []OrderDto `json:"order_dtos,omitempty" xml:"order_dtos>order_dto,omitempty"`
    // 每页的数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 当前页的数量
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
    // 当前页
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 是否有下一页
    HasNextPage   bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
    // 总页数
    TotalPageNum   int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
    // 结果总数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // 返回结果
    ProviderMemberVo   *ProviderMemberVo `json:"provider_member_vo,omitempty" xml:"provider_member_vo,omitempty"`
}
