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
    _bizData   string
    // 进度列表
    _progressList   []string
    // 附件列表
    _claimAttachments   []ClaimAttachment
    // 保单业务单号
    _policyBizNo   string
    // 外部业务单号
    _outBizNo   string
    // 业务来源
    _bizSource   string
    // 理赔金额(单位为分)
    _claimFee   int64
    // 理赔单号
    _claimNo   string
    // 理赔外部业务单号
    _claimOutBizNo   string
    // 标准产品ID
    _spNo   string
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
func (r *AlipayBaoxianClaimUpdateRequest) SetBizData(_bizData string) error {
    r._bizData = _bizData
    r.Set("biz_data", _bizData)
    return nil
}

// BizData Getter
func (r AlipayBaoxianClaimUpdateRequest) GetBizData() string {
    return r._bizData
}
// ProgressList Setter
// 进度列表
func (r *AlipayBaoxianClaimUpdateRequest) SetProgressList(_progressList []string) error {
    r._progressList = _progressList
    r.Set("progress_list", _progressList)
    return nil
}

// ProgressList Getter
func (r AlipayBaoxianClaimUpdateRequest) GetProgressList() []string {
    return r._progressList
}
// ClaimAttachments Setter
// 附件列表
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimAttachments(_claimAttachments []ClaimAttachment) error {
    r._claimAttachments = _claimAttachments
    r.Set("claim_attachments", _claimAttachments)
    return nil
}

// ClaimAttachments Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimAttachments() []ClaimAttachment {
    return r._claimAttachments
}
// PolicyBizNo Setter
// 保单业务单号
func (r *AlipayBaoxianClaimUpdateRequest) SetPolicyBizNo(_policyBizNo string) error {
    r._policyBizNo = _policyBizNo
    r.Set("policy_biz_no", _policyBizNo)
    return nil
}

// PolicyBizNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetPolicyBizNo() string {
    return r._policyBizNo
}
// OutBizNo Setter
// 外部业务单号
func (r *AlipayBaoxianClaimUpdateRequest) SetOutBizNo(_outBizNo string) error {
    r._outBizNo = _outBizNo
    r.Set("out_biz_no", _outBizNo)
    return nil
}

// OutBizNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetOutBizNo() string {
    return r._outBizNo
}
// BizSource Setter
// 业务来源
func (r *AlipayBaoxianClaimUpdateRequest) SetBizSource(_bizSource string) error {
    r._bizSource = _bizSource
    r.Set("biz_source", _bizSource)
    return nil
}

// BizSource Getter
func (r AlipayBaoxianClaimUpdateRequest) GetBizSource() string {
    return r._bizSource
}
// ClaimFee Setter
// 理赔金额(单位为分)
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimFee(_claimFee int64) error {
    r._claimFee = _claimFee
    r.Set("claim_fee", _claimFee)
    return nil
}

// ClaimFee Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimFee() int64 {
    return r._claimFee
}
// ClaimNo Setter
// 理赔单号
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimNo(_claimNo string) error {
    r._claimNo = _claimNo
    r.Set("claim_no", _claimNo)
    return nil
}

// ClaimNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimNo() string {
    return r._claimNo
}
// ClaimOutBizNo Setter
// 理赔外部业务单号
func (r *AlipayBaoxianClaimUpdateRequest) SetClaimOutBizNo(_claimOutBizNo string) error {
    r._claimOutBizNo = _claimOutBizNo
    r.Set("claim_out_biz_no", _claimOutBizNo)
    return nil
}

// ClaimOutBizNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetClaimOutBizNo() string {
    return r._claimOutBizNo
}
// SpNo Setter
// 标准产品ID
func (r *AlipayBaoxianClaimUpdateRequest) SetSpNo(_spNo string) error {
    r._spNo = _spNo
    r.Set("sp_no", _spNo)
    return nil
}

// SpNo Getter
func (r AlipayBaoxianClaimUpdateRequest) GetSpNo() string {
    return r._spNo
}
