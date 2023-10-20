package baoxian

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlipayBaoxianClaimUpdateAPIRequest 更新赔案 API请求
// alipay.baoxian.claim.update
//
// 更新保险理赔单
type AlipayBaoxianClaimUpdateAPIRequest struct {
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

// NewAlipayBaoxianClaimUpdateRequest 初始化AlipayBaoxianClaimUpdateAPIRequest对象
func NewAlipayBaoxianClaimUpdateRequest() *AlipayBaoxianClaimUpdateAPIRequest {
	return &AlipayBaoxianClaimUpdateAPIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlipayBaoxianClaimUpdateAPIRequest) Reset() {
	r._progressList = r._progressList[:0]
	r._claimAttachments = r._claimAttachments[:0]
	r._bizData = ""
	r._policyBizNo = ""
	r._outBizNo = ""
	r._bizSource = ""
	r._claimNo = ""
	r._claimOutBizNo = ""
	r._spNo = ""
	r._claimFee = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlipayBaoxianClaimUpdateAPIRequest) GetApiMethodName() string {
	return "alipay.baoxian.claim.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlipayBaoxianClaimUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlipayBaoxianClaimUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProgressList is ProgressList Setter
// 进度列表
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetProgressList(_progressList []string) error {
	r._progressList = _progressList
	r.Set("progress_list", _progressList)
	return nil
}

// GetProgressList ProgressList Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetProgressList() []string {
	return r._progressList
}

// SetClaimAttachments is ClaimAttachments Setter
// 附件列表
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetClaimAttachments(_claimAttachments []ClaimAttachment) error {
	r._claimAttachments = _claimAttachments
	r.Set("claim_attachments", _claimAttachments)
	return nil
}

// GetClaimAttachments ClaimAttachments Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetClaimAttachments() []ClaimAttachment {
	return r._claimAttachments
}

// SetBizData is BizData Setter
// 业务数据
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetBizData(_bizData string) error {
	r._bizData = _bizData
	r.Set("biz_data", _bizData)
	return nil
}

// GetBizData BizData Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetBizData() string {
	return r._bizData
}

// SetPolicyBizNo is PolicyBizNo Setter
// 保单业务单号
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetPolicyBizNo(_policyBizNo string) error {
	r._policyBizNo = _policyBizNo
	r.Set("policy_biz_no", _policyBizNo)
	return nil
}

// GetPolicyBizNo PolicyBizNo Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetPolicyBizNo() string {
	return r._policyBizNo
}

// SetOutBizNo is OutBizNo Setter
// 外部业务单号
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetOutBizNo(_outBizNo string) error {
	r._outBizNo = _outBizNo
	r.Set("out_biz_no", _outBizNo)
	return nil
}

// GetOutBizNo OutBizNo Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetOutBizNo() string {
	return r._outBizNo
}

// SetBizSource is BizSource Setter
// 业务来源
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetBizSource(_bizSource string) error {
	r._bizSource = _bizSource
	r.Set("biz_source", _bizSource)
	return nil
}

// GetBizSource BizSource Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetBizSource() string {
	return r._bizSource
}

// SetClaimNo is ClaimNo Setter
// 理赔单号
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetClaimNo(_claimNo string) error {
	r._claimNo = _claimNo
	r.Set("claim_no", _claimNo)
	return nil
}

// GetClaimNo ClaimNo Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetClaimNo() string {
	return r._claimNo
}

// SetClaimOutBizNo is ClaimOutBizNo Setter
// 理赔外部业务单号
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetClaimOutBizNo(_claimOutBizNo string) error {
	r._claimOutBizNo = _claimOutBizNo
	r.Set("claim_out_biz_no", _claimOutBizNo)
	return nil
}

// GetClaimOutBizNo ClaimOutBizNo Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetClaimOutBizNo() string {
	return r._claimOutBizNo
}

// SetSpNo is SpNo Setter
// 标准产品ID
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetSpNo(_spNo string) error {
	r._spNo = _spNo
	r.Set("sp_no", _spNo)
	return nil
}

// GetSpNo SpNo Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetSpNo() string {
	return r._spNo
}

// SetClaimFee is ClaimFee Setter
// 理赔金额(单位为分)
func (r *AlipayBaoxianClaimUpdateAPIRequest) SetClaimFee(_claimFee int64) error {
	r._claimFee = _claimFee
	r.Set("claim_fee", _claimFee)
	return nil
}

// GetClaimFee ClaimFee Getter
func (r AlipayBaoxianClaimUpdateAPIRequest) GetClaimFee() int64 {
	return r._claimFee
}

var poolAlipayBaoxianClaimUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlipayBaoxianClaimUpdateRequest()
	},
}

// GetAlipayBaoxianClaimUpdateRequest 从 sync.Pool 获取 AlipayBaoxianClaimUpdateAPIRequest
func GetAlipayBaoxianClaimUpdateAPIRequest() *AlipayBaoxianClaimUpdateAPIRequest {
	return poolAlipayBaoxianClaimUpdateAPIRequest.Get().(*AlipayBaoxianClaimUpdateAPIRequest)
}

// ReleaseAlipayBaoxianClaimUpdateAPIRequest 将 AlipayBaoxianClaimUpdateAPIRequest 放入 sync.Pool
func ReleaseAlipayBaoxianClaimUpdateAPIRequest(v *AlipayBaoxianClaimUpdateAPIRequest) {
	v.Reset()
	poolAlipayBaoxianClaimUpdateAPIRequest.Put(v)
}
