package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbreviewaddAPIRequest 对外开放评论接口 API请求
// taobao.xhotel.bnbreview.add
//
// 对外开放评论接口
type TaobaoxhotelbnbreviewaddAPIRequest struct {
	model.Params
	// 入住时间
	_checkInTime string
	// 创建时间
	_gmtCreate string
	// 评论内容
	_content string
	// 用户名称
	_userNick string
	// 总评分,Double类型得
	_totalScore string
	// 评论来源ID，38-小猪 39-爱彼迎
	_source int64
	// 图片地址
	_picInfoList *ReviewPicInfo
	// 飞猪侧房源ID
	_rid int64
	// 外部评论id
	_outerId int64
	// 评分细分
	_scoreDetail *ReviewDetailInfo
}

// NewTaobaoxhotelbnbreviewaddRequest 初始化TaobaoxhotelbnbreviewaddAPIRequest对象
func NewTaobaoxhotelbnbreviewaddRequest() *TaobaoxhotelbnbreviewaddAPIRequest {
	return &TaobaoxhotelbnbreviewaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbreview.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckInTime is CheckInTime Setter
// 入住时间
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetCheckInTime(_checkInTime string) error {
	r._checkInTime = _checkInTime
	r.Set("check_in_time", _checkInTime)
	return nil
}

// GetCheckInTime CheckInTime Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetCheckInTime() string {
	return r._checkInTime
}

// SetGmtCreate is GmtCreate Setter
// 创建时间
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetGmtCreate(_gmtCreate string) error {
	r._gmtCreate = _gmtCreate
	r.Set("gmt_create", _gmtCreate)
	return nil
}

// GetGmtCreate GmtCreate Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetGmtCreate() string {
	return r._gmtCreate
}

// SetContent is Content Setter
// 评论内容
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetContent() string {
	return r._content
}

// SetUserNick is UserNick Setter
// 用户名称
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetTotalScore is TotalScore Setter
// 总评分,Double类型得
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetTotalScore(_totalScore string) error {
	r._totalScore = _totalScore
	r.Set("total_score", _totalScore)
	return nil
}

// GetTotalScore TotalScore Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetTotalScore() string {
	return r._totalScore
}

// SetSource is Source Setter
// 评论来源ID，38-小猪 39-爱彼迎
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetSource(_source int64) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetSource() int64 {
	return r._source
}

// SetPicInfoList is PicInfoList Setter
// 图片地址
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetPicInfoList(_picInfoList *ReviewPicInfo) error {
	r._picInfoList = _picInfoList
	r.Set("pic_info_list", _picInfoList)
	return nil
}

// GetPicInfoList PicInfoList Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetPicInfoList() *ReviewPicInfo {
	return r._picInfoList
}

// SetRid is Rid Setter
// 飞猪侧房源ID
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetRid() int64 {
	return r._rid
}

// SetOuterId is OuterId Setter
// 外部评论id
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetOuterId(_outerId int64) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetOuterId() int64 {
	return r._outerId
}

// SetScoreDetail is ScoreDetail Setter
// 评分细分
func (r *TaobaoxhotelbnbreviewaddAPIRequest) SetScoreDetail(_scoreDetail *ReviewDetailInfo) error {
	r._scoreDetail = _scoreDetail
	r.Set("score_detail", _scoreDetail)
	return nil
}

// GetScoreDetail ScoreDetail Getter
func (r TaobaoxhotelbnbreviewaddAPIRequest) GetScoreDetail() *ReviewDetailInfo {
	return r._scoreDetail
}
