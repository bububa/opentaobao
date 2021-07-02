package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsDataSportsSyncmatchdataAPIRequest 阿里体育数据中心用户比赛数据同步接口 API请求
// alibaba.alisports.data.sports.syncmatchdata
//
// 阿里体育数据中心用户比赛数据同步接口
type AlibabaAlisportsDataSportsSyncmatchdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 成绩(比赛用时)
	_score string
	// 组别 1001半程马拉松  1002全程马拉松
	_matchGroup int64
	// 国家
	_country string
	// 出生日期 格式：Y-m-d
	_birthday string
	// 手机号
	_mobile string
	// 证件ID
	_cardId string
	// 证件类型 1身份证 2军官证 4护照 8台胞证 16港澳通行证 32未设置  64 其他
	_cardType int64
	// 类型：1专业 2业余
	_type int64
	// 性别 0未知 1男 2女
	_gender int64
	// 姓名
	_name string
	// 排名
	_ranking int64
	// 比赛名（展示用）
	_match string
	// 比赛类型 1马拉松
	_matchType int64
	// 参赛号
	_num string
	// 阿里体育用户id
	_aliuid string
	// 接口签名
	_alispSign string
	// 时间戳精确到秒
	_alispTime string
	// 比赛日期 格式：Y-m-d
	_matchTime string
	// 枪声成绩
	_gunshotScore string
	// 枪声排名
	_gunshotRanking int64
	// 平均配速
	_speed string
	// 展示用，比如：男子半程马拉松
	_project string
	// 比如马拉松 5KM 10KM 15KM分段成绩，json key->value 格式
	_subScore string
	// 比如马拉松 5KM 10KM 15KM分段时间点，json key->value 格式
	_subPoint string
}

// NewAlibabaAlisportsDataSportsSyncmatchdataRequest 初始化AlibabaAlisportsDataSportsSyncmatchdataAPIRequest对象
func NewAlibabaAlisportsDataSportsSyncmatchdataRequest() *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest {
	return &AlibabaAlisportsDataSportsSyncmatchdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.data.sports.syncmatchdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAlispAppKey is AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetScore is Score Setter
// 成绩(比赛用时)
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetScore(_score string) error {
	r._score = _score
	r.Set("score", _score)
	return nil
}

// GetScore Score Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetScore() string {
	return r._score
}

// SetMatchGroup is MatchGroup Setter
// 组别 1001半程马拉松  1002全程马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetMatchGroup(_matchGroup int64) error {
	r._matchGroup = _matchGroup
	r.Set("match_group", _matchGroup)
	return nil
}

// GetMatchGroup MatchGroup Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetMatchGroup() int64 {
	return r._matchGroup
}

// SetCountry is Country Setter
// 国家
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetCountry() string {
	return r._country
}

// SetBirthday is Birthday Setter
// 出生日期 格式：Y-m-d
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetBirthday(_birthday string) error {
	r._birthday = _birthday
	r.Set("birthday", _birthday)
	return nil
}

// GetBirthday Birthday Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetBirthday() string {
	return r._birthday
}

// SetMobile is Mobile Setter
// 手机号
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetMobile() string {
	return r._mobile
}

// SetCardId is CardId Setter
// 证件ID
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetCardId(_cardId string) error {
	r._cardId = _cardId
	r.Set("card_id", _cardId)
	return nil
}

// GetCardId CardId Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetCardId() string {
	return r._cardId
}

// SetCardType is CardType Setter
// 证件类型 1身份证 2军官证 4护照 8台胞证 16港澳通行证 32未设置  64 其他
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetCardType(_cardType int64) error {
	r._cardType = _cardType
	r.Set("card_type", _cardType)
	return nil
}

// GetCardType CardType Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetCardType() int64 {
	return r._cardType
}

// SetType is Type Setter
// 类型：1专业 2业余
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetType() int64 {
	return r._type
}

// SetGender is Gender Setter
// 性别 0未知 1男 2女
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetGender(_gender int64) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetGender() int64 {
	return r._gender
}

// SetName is Name Setter
// 姓名
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetName() string {
	return r._name
}

// SetRanking is Ranking Setter
// 排名
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetRanking(_ranking int64) error {
	r._ranking = _ranking
	r.Set("ranking", _ranking)
	return nil
}

// GetRanking Ranking Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetRanking() int64 {
	return r._ranking
}

// SetMatch is Match Setter
// 比赛名（展示用）
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetMatch(_match string) error {
	r._match = _match
	r.Set("match", _match)
	return nil
}

// GetMatch Match Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetMatch() string {
	return r._match
}

// SetMatchType is MatchType Setter
// 比赛类型 1马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetMatchType(_matchType int64) error {
	r._matchType = _matchType
	r.Set("match_type", _matchType)
	return nil
}

// GetMatchType MatchType Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetMatchType() int64 {
	return r._matchType
}

// SetNum is Num Setter
// 参赛号
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetNum(_num string) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// GetNum Num Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetNum() string {
	return r._num
}

// SetAliuid is Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetAlispSign is AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAlispTime is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetMatchTime is MatchTime Setter
// 比赛日期 格式：Y-m-d
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetMatchTime(_matchTime string) error {
	r._matchTime = _matchTime
	r.Set("match_time", _matchTime)
	return nil
}

// GetMatchTime MatchTime Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetMatchTime() string {
	return r._matchTime
}

// SetGunshotScore is GunshotScore Setter
// 枪声成绩
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetGunshotScore(_gunshotScore string) error {
	r._gunshotScore = _gunshotScore
	r.Set("gunshot_score", _gunshotScore)
	return nil
}

// GetGunshotScore GunshotScore Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetGunshotScore() string {
	return r._gunshotScore
}

// SetGunshotRanking is GunshotRanking Setter
// 枪声排名
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetGunshotRanking(_gunshotRanking int64) error {
	r._gunshotRanking = _gunshotRanking
	r.Set("gunshot_ranking", _gunshotRanking)
	return nil
}

// GetGunshotRanking GunshotRanking Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetGunshotRanking() int64 {
	return r._gunshotRanking
}

// SetSpeed is Speed Setter
// 平均配速
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetSpeed(_speed string) error {
	r._speed = _speed
	r.Set("speed", _speed)
	return nil
}

// GetSpeed Speed Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetSpeed() string {
	return r._speed
}

// SetProject is Project Setter
// 展示用，比如：男子半程马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetProject(_project string) error {
	r._project = _project
	r.Set("project", _project)
	return nil
}

// GetProject Project Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetProject() string {
	return r._project
}

// SetSubScore is SubScore Setter
// 比如马拉松 5KM 10KM 15KM分段成绩，json key->value 格式
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetSubScore(_subScore string) error {
	r._subScore = _subScore
	r.Set("sub_score", _subScore)
	return nil
}

// GetSubScore SubScore Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetSubScore() string {
	return r._subScore
}

// SetSubPoint is SubPoint Setter
// 比如马拉松 5KM 10KM 15KM分段时间点，json key->value 格式
func (r *AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) SetSubPoint(_subPoint string) error {
	r._subPoint = _subPoint
	r.Set("sub_point", _subPoint)
	return nil
}

// GetSubPoint SubPoint Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataAPIRequest) GetSubPoint() string {
	return r._subPoint
}
