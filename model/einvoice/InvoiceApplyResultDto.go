package einvoice

// InvoiceApplyResultDto 
type InvoiceApplyResultDto struct {
    // 中台发票申请ID，由中台生成。字母或数字组成
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 申请状态，可选值：  applying: 申请中，初始状态；  cancelled: 申请已取消；  confirmed: 商户已确认，待开/录入发票；  craeting_inv: 开票中，待发票结果回传；  inv_failed: 开票失败；  inv_success: 开票成功；  inv_part_success: 部分成功（拆单场景下存在。举例：发票申请拆单之后有10张票，其中有1张开票成功时，此时申请状态即为inv_part_success，当10张票全部成功申请状态则为inv_success）
    ApplyStatus   string `json:"apply_status,omitempty" xml:"apply_status,omitempty"`
    // 生成的发票申请页面URL, 用户可在该页面中填写抬头等信息，然后提交正式的发票申请。  当apply_mode=create_apply_url 时必须返回。
    ApplyUrl   string `json:"apply_url,omitempty" xml:"apply_url,omitempty"`
    // 开票结果
    CreateInvResultList   []InvoiceCreateSimpleResultDto `json:"create_inv_result_list,omitempty" xml:"create_inv_result_list>invoice_create_simple_result_dto,omitempty"`
}
