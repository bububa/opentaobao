package baoxian

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新赔案 APIRequest
alipay.baoxian.claim.update

更新保险理赔单
*/
type AlipayBaoxianClaimUpdateRequest struct {
    model.Params

    // 业务数据
    bizData   string 

    // 进度列表
    progressList   []Json 

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

func NewAlipayBaoxianClaimUpdateRequest() *AlipayBaoxianClaimUpdateRequest{
    return &AlipayBaoxianClaimUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlipayBaoxianClaimUpdateRequest) GetApiMethodName() string {
    return "alipay.baoxian.claim.update"
}

func (r AlipayBaoxianClaimUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlipayBaoxianClaimUpdateRequest) SetBizData(bizData string) error {
    r.bizData = bizData
    r.Set("biz_data", bizData)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetBizData() string {
    return r.bizData
}

func (r *AlipayBaoxianClaimUpdateRequest) SetProgressList(progressList []Json) error {
    r.progressList = progressList
    r.Set("progress_list", progressList)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetProgressList() []Json {
    return r.progressList
}

func (r *AlipayBaoxianClaimUpdateRequest) SetClaimAttachments(claimAttachments []ClaimAttachment) error {
    r.claimAttachments = claimAttachments
    r.Set("claim_attachments", claimAttachments)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetClaimAttachments() []ClaimAttachment {
    return r.claimAttachments
}

func (r *AlipayBaoxianClaimUpdateRequest) SetPolicyBizNo(policyBizNo string) error {
    r.policyBizNo = policyBizNo
    r.Set("policy_biz_no", policyBizNo)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetPolicyBizNo() string {
    return r.policyBizNo
}

func (r *AlipayBaoxianClaimUpdateRequest) SetOutBizNo(outBizNo string) error {
    r.outBizNo = outBizNo
    r.Set("out_biz_no", outBizNo)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetOutBizNo() string {
    return r.outBizNo
}

func (r *AlipayBaoxianClaimUpdateRequest) SetBizSource(bizSource string) error {
    r.bizSource = bizSource
    r.Set("biz_source", bizSource)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetBizSource() string {
    return r.bizSource
}

func (r *AlipayBaoxianClaimUpdateRequest) SetClaimFee(claimFee int64) error {
    r.claimFee = claimFee
    r.Set("claim_fee", claimFee)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetClaimFee() int64 {
    return r.claimFee
}

func (r *AlipayBaoxianClaimUpdateRequest) SetClaimNo(claimNo string) error {
    r.claimNo = claimNo
    r.Set("claim_no", claimNo)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetClaimNo() string {
    return r.claimNo
}

func (r *AlipayBaoxianClaimUpdateRequest) SetClaimOutBizNo(claimOutBizNo string) error {
    r.claimOutBizNo = claimOutBizNo
    r.Set("claim_out_biz_no", claimOutBizNo)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetClaimOutBizNo() string {
    return r.claimOutBizNo
}

func (r *AlipayBaoxianClaimUpdateRequest) SetSpNo(spNo string) error {
    r.spNo = spNo
    r.Set("sp_no", spNo)
    return nil
}

func (r AlipayBaoxianClaimUpdateRequest) GetSpNo() string {
    return r.spNo
}

