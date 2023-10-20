package tmic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltmicquestionnairesurveygetAPIRequest 天猫新品创新中心问卷数据获取 API请求
// tmall.tmic.questionnaire.survey.get
//
// 天猫新品创新中心问卷数据获取
type TmalltmicquestionnairesurveygetAPIRequest struct {
	model.Params
	// 问卷hashCode，问卷对外唯一编码
	_hashCode string
	// biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
	_biz string
	// open_id
	_openUserId string
}

// NewTmalltmicquestionnairesurveygetRequest 初始化TmalltmicquestionnairesurveygetAPIRequest对象
func NewTmalltmicquestionnairesurveygetRequest() *TmalltmicquestionnairesurveygetAPIRequest {
	return &TmalltmicquestionnairesurveygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltmicquestionnairesurveygetAPIRequest) GetApiMethodName() string {
	return "tmall.tmic.questionnaire.survey.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltmicquestionnairesurveygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltmicquestionnairesurveygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHashCode is HashCode Setter
// 问卷hashCode，问卷对外唯一编码
func (r *TmalltmicquestionnairesurveygetAPIRequest) SetHashCode(_hashCode string) error {
	r._hashCode = _hashCode
	r.Set("hash_code", _hashCode)
	return nil
}

// GetHashCode HashCode Getter
func (r TmalltmicquestionnairesurveygetAPIRequest) GetHashCode() string {
	return r._hashCode
}

// SetBiz is Biz Setter
// biz业务参数，1024表示投放id，下划线分隔，fav表示收藏行为的英文
func (r *TmalltmicquestionnairesurveygetAPIRequest) SetBiz(_biz string) error {
	r._biz = _biz
	r.Set("biz", _biz)
	return nil
}

// GetBiz Biz Getter
func (r TmalltmicquestionnairesurveygetAPIRequest) GetBiz() string {
	return r._biz
}

// SetOpenUserId is OpenUserId Setter
// open_id
func (r *TmalltmicquestionnairesurveygetAPIRequest) SetOpenUserId(_openUserId string) error {
	r._openUserId = _openUserId
	r.Set("open_user_id", _openUserId)
	return nil
}

// GetOpenUserId OpenUserId Getter
func (r TmalltmicquestionnairesurveygetAPIRequest) GetOpenUserId() string {
	return r._openUserId
}
