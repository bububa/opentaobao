package baoxian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlipaybaoxianclaimupdateAPIRequest 更新赔案 API请求
// alipay.baoxian.claim.update
//
// 更新保险理赔单
type AlipaybaoxianclaimupdateAPIRequest struct {
	model.Params
	// 进度列表
	_progressList []string
	// 附件列表
	_claimAttachments []ClaimAttachment
	// 业务数据
	_bizData string
	// 保单业务单号
	_policyBizNo string
	// 外部业务单号
	_outBizNo string
	// 业务来源
	_bizSource string
	// 理赔单号
	_claimNo string
	// 理赔外部业务单号
	_claimOutBizNo string
	// 标准产品ID
	_spNo string
	// 理赔金额(单位为分)
	_claimFee int64
}

// NewAlipaybaoxianclaimupdateRequest 初始化AlipaybaoxianclaimupdateAPIRequest对象
func NewAlipaybaoxianclaimupdateRequest() *AlipaybaoxianclaimupdateAPIRequest {
	return &AlipaybaoxianclaimupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipaybaoxianclaimupdateAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipaybaoxianclaimupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipaybaoxianclaimupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProgressList is ProgressList Setter
// 进度列表
func (r *AlipaybaoxianclaimupdateAPIRequest) SetProgressList(_progressList []string) error {
	r._progressList = _progressList
	r.Set("progress_list", _progressList)
	return nil
}

// GetProgressList ProgressList Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetProgressList() []string {
	return r._progressList
}

// SetClaimAttachments is ClaimAttachments Setter
// 附件列表
func (r *AlipaybaoxianclaimupdateAPIRequest) SetClaimAttachments(_claimAttachments []ClaimAttachment) error {
	r._claimAttachments = _claimAttachments
	r.Set("claim_attachments", _claimAttachments)
	return nil
}

// GetClaimAttachments ClaimAttachments Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetClaimAttachments() []ClaimAttachment {
	return r._claimAttachments
}

// SetBizData is BizData Setter
// 业务数据
func (r *AlipaybaoxianclaimupdateAPIRequest) SetBizData(_bizData string) error {
	r._bizData = _bizData
	r.Set("biz_data", _bizData)
	return nil
}

// GetBizData BizData Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetBizData() string {
	return r._bizData
}

// SetPolicyBizNo is PolicyBizNo Setter
// 保单业务单号
func (r *AlipaybaoxianclaimupdateAPIRequest) SetPolicyBizNo(_policyBizNo string) error {
	r._policyBizNo = _policyBizNo
	r.Set("policy_biz_no", _policyBizNo)
	return nil
}

// GetPolicyBizNo PolicyBizNo Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetPolicyBizNo() string {
	return r._policyBizNo
}

// SetOutBizNo is OutBizNo Setter
// 外部业务单号
func (r *AlipaybaoxianclaimupdateAPIRequest) SetOutBizNo(_outBizNo string) error {
	r._outBizNo = _outBizNo
	r.Set("out_biz_no", _outBizNo)
	return nil
}

// GetOutBizNo OutBizNo Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetOutBizNo() string {
	return r._outBizNo
}

// SetBizSource is BizSource Setter
// 业务来源
func (r *AlipaybaoxianclaimupdateAPIRequest) SetBizSource(_bizSource string) error {
	r._bizSource = _bizSource
	r.Set("biz_source", _bizSource)
	return nil
}

// GetBizSource BizSource Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetBizSource() string {
	return r._bizSource
}

// SetClaimNo is ClaimNo Setter
// 理赔单号
func (r *AlipaybaoxianclaimupdateAPIRequest) SetClaimNo(_claimNo string) error {
	r._claimNo = _claimNo
	r.Set("claim_no", _claimNo)
	return nil
}

// GetClaimNo ClaimNo Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetClaimNo() string {
	return r._claimNo
}

// SetClaimOutBizNo is ClaimOutBizNo Setter
// 理赔外部业务单号
func (r *AlipaybaoxianclaimupdateAPIRequest) SetClaimOutBizNo(_claimOutBizNo string) error {
	r._claimOutBizNo = _claimOutBizNo
	r.Set("claim_out_biz_no", _claimOutBizNo)
	return nil
}

// GetClaimOutBizNo ClaimOutBizNo Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetClaimOutBizNo() string {
	return r._claimOutBizNo
}

// SetSpNo is SpNo Setter
// 标准产品ID
func (r *AlipaybaoxianclaimupdateAPIRequest) SetSpNo(_spNo string) error {
	r._spNo = _spNo
	r.Set("sp_no", _spNo)
	return nil
}

// GetSpNo SpNo Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetSpNo() string {
	return r._spNo
}

// SetClaimFee is ClaimFee Setter
// 理赔金额(单位为分)
func (r *AlipaybaoxianclaimupdateAPIRequest) SetClaimFee(_claimFee int64) error {
	r._claimFee = _claimFee
	r.Set("claim_fee", _claimFee)
	return nil
}

// GetClaimFee ClaimFee Getter
func (r AlipaybaoxianclaimupdateAPIRequest) GetClaimFee() int64 {
	return r._claimFee
}
