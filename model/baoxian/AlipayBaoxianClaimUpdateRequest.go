package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新赔案 API请求
alipay.baoxian.claim.update

更新保险理赔单
*/
type AlipayBaoxianClaimUpdateRequest struct {
    model.Params
    // 业务数据
    bizData   string
    // 进度列表
    progressList   []string
    // 附件列表
    claimAttachments   []ClaimAttachment
    // 保单业务单号
    policyBizNo   string
    // 外部业务单号
    outBizNo   string
    // 业务来源
    bizSource   string
    // 理赔金额(单位为分)
    claimFee   int64
    // 理赔单号
    claimNo   string
    // 理赔外部业务单号
    claimOutBizNo   string
    // 标准产品ID
    spNo   string
}

// 初始化AlipayBaoxianClaimUpdateRequest对象
func NewAlipayBaoxianClaimUpdateRequest() *AlipayBaoxianClaimUpdateRequest{
    return &AlipayBaoxianClaimUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimUpdateRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.update"
}

// IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizData Setter
// 业务数据
func (r *AlipayBaoxianClaimUpdateRequest) SetBizData(bizData string) error {
    r.bizData = bizData
    r.Set("biz_data", bizData)
    return nil
}

// BizData Getter
func (r AlipayBaoxianClaimUpdateRequest) GetBizData() string {
    return r.bizData
}
// ProgressList Setter
// 进度列表
func (r *AlipayBaoxianClaimUpdateRequest) SetProgressList(progressList []string) error {
    r.progressList = progressList
    r.Set("progress_list", progressList)
    return nil
}

// ProgressList Getter
func (r AlipayBaoxianClaimUpdateRequest) GetProgressList() []string {
    return r.progressList
}
// ClaimAttachments Setter
// 附件列表
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimAttachments(claimAttachments []ClaimAttachment) error {
    r.claimAttachments = claimAttachments
    r.Set("claim_attachments", claimAttachments)
    return nil
}

// ClaimAttachments Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimAttachments() []ClaimAttachment {
    return r.claimAttachments
}
// PolicyBizNo Setter
// 保单业务单号
func (r *AlipayBaoxianClaimUpdateRequest) SetPolicyBizNo(policyBizNo string) error {
    r.policyBizNo = policyBizNo
    r.Set("policy_biz_no", policyBizNo)
    return nil
}

// PolicyBizNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetPolicyBizNo() string {
    return r.policyBizNo
}
// OutBizNo Setter
// 外部业务单号
func (r *AlipayBaoxianClaimUpdateRequest) SetOutBizNo(outBizNo string) error {
    r.outBizNo = outBizNo
    r.Set("out_biz_no", outBizNo)
    return nil
}

// OutBizNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetOutBizNo() string {
    return r.outBizNo
}
// BizSource Setter
// 业务来源
func (r *AlipayBaoxianClaimUpdateRequest) SetBizSource(bizSource string) error {
    r.bizSource = bizSource
    r.Set("biz_source", bizSource)
    return nil
}

// BizSource Getter
func (r AlipayBaoxianClaimUpdateRequest) GetBizSource() string {
    return r.bizSource
}
// ClaimFee Setter
// 理赔金额(单位为分)
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimFee(claimFee int64) error {
    r.claimFee = claimFee
    r.Set("claim_fee", claimFee)
    return nil
}

// ClaimFee Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimFee() int64 {
    return r.claimFee
}
// ClaimNo Setter
// 理赔单号
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimNo(claimNo string) error {
    r.claimNo = claimNo
    r.Set("claim_no", claimNo)
    return nil
}

// ClaimNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimNo() string {
    return r.claimNo
}
// ClaimOutBizNo Setter
// 理赔外部业务单号
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimOutBizNo(claimOutBizNo string) error {
    r.claimOutBizNo = claimOutBizNo
    r.Set("claim_out_biz_no", claimOutBizNo)
    return nil
}

// ClaimOutBizNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimOutBizNo() string {
    return r.claimOutBizNo
}
// SpNo Setter
// 标准产品ID
func (r *AlipayBaoxianClaimUpdateRequest) SetSpNo(spNo string) error {
    r.spNo = spNo
    r.Set("sp_no", spNo)
    return nil
}

// SpNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetSpNo() string {
    return r.spNo
}
