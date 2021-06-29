package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户比赛数据同步接口 API请求
alibaba.alisports.data.sports.syncmatchdata

阿里体育数据中心用户比赛数据同步接口
*/
type AlibabaAlisportsDataSportsSyncmatchdataRequest struct {
    model.Params
    // 应用appkey
    _alispAppKey   string
    // 成绩(比赛用时)
    _score   string
    // 组别 1001半程马拉松  1002全程马拉松
    _matchGroup   int64
    // 国家
    _country   string
    // 出生日期 格式：Y-m-d
    _birthday   string
    // 手机号
    _mobile   string
    // 证件ID
    _cardId   string
    // 证件类型 1身份证 2军官证 4护照 8台胞证 16港澳通行证 32未设置  64 其他
    _cardType   int64
    // 类型：1专业 2业余
    _type   int64
    // 性别 0未知 1男 2女
    _gender   int64
    // 姓名
    _name   string
    // 排名
    _ranking   int64
    // 比赛名（展示用）
    _match   string
    // 比赛类型 1马拉松
    _matchType   int64
    // 参赛号
    _num   string
    // 阿里体育用户id
    _aliuid   string
    // 接口签名
    _alispSign   string
    // 时间戳精确到秒
    _alispTime   string
    // 比赛日期 格式：Y-m-d
    _matchTime   string
    // 枪声成绩
    _gunshotScore   string
    // 枪声排名
    _gunshotRanking   int64
    // 平均配速
    _speed   string
    // 展示用，比如：男子半程马拉松
    _project   string
    // 比如马拉松 5KM 10KM 15KM分段成绩，json key->value 格式
    _subScore   string
    // 比如马拉松 5KM 10KM 15KM分段时间点，json key->value 格式
    _subPoint   string
}

// 初始化AlibabaAlisportsDataSportsSyncmatchdataRequest对象
func NewAlibabaAlisportsDataSportsSyncmatchdataRequest() *AlibabaAlisportsDataSportsSyncmatchdataRequest{
    return &AlibabaAlisportsDataSportsSyncmatchdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncmatchdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlispAppKey Setter
// 应用appkey
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// Score Setter
// 成绩(比赛用时)
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetScore(_score string) error {
    r._score = _score
    r.Set("score", _score)
    return nil
}

// Score Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetScore() string {
    return r._score
}
// MatchGroup Setter
// 组别 1001半程马拉松  1002全程马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchGroup(_matchGroup int64) error {
    r._matchGroup = _matchGroup
    r.Set("match_group", _matchGroup)
    return nil
}

// MatchGroup Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchGroup() int64 {
    return r._matchGroup
}
// Country Setter
// 国家
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCountry(_country string) error {
    r._country = _country
    r.Set("country", _country)
    return nil
}

// Country Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCountry() string {
    return r._country
}
// Birthday Setter
// 出生日期 格式：Y-m-d
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetBirthday(_birthday string) error {
    r._birthday = _birthday
    r.Set("birthday", _birthday)
    return nil
}

// Birthday Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetBirthday() string {
    return r._birthday
}
// Mobile Setter
// 手机号
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMobile() string {
    return r._mobile
}
// CardId Setter
// 证件ID
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCardId(_cardId string) error {
    r._cardId = _cardId
    r.Set("card_id", _cardId)
    return nil
}

// CardId Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCardId() string {
    return r._cardId
}
// CardType Setter
// 证件类型 1身份证 2军官证 4护照 8台胞证 16港澳通行证 32未设置  64 其他
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCardType(_cardType int64) error {
    r._cardType = _cardType
    r.Set("card_type", _cardType)
    return nil
}

// CardType Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCardType() int64 {
    return r._cardType
}
// Type Setter
// 类型：1专业 2业余
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetType() int64 {
    return r._type
}
// Gender Setter
// 性别 0未知 1男 2女
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGender(_gender int64) error {
    r._gender = _gender
    r.Set("gender", _gender)
    return nil
}

// Gender Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGender() int64 {
    return r._gender
}
// Name Setter
// 姓名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetName() string {
    return r._name
}
// Ranking Setter
// 排名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetRanking(_ranking int64) error {
    r._ranking = _ranking
    r.Set("ranking", _ranking)
    return nil
}

// Ranking Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetRanking() int64 {
    return r._ranking
}
// Match Setter
// 比赛名（展示用）
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatch(_match string) error {
    r._match = _match
    r.Set("match", _match)
    return nil
}

// Match Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatch() string {
    return r._match
}
// MatchType Setter
// 比赛类型 1马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchType(_matchType int64) error {
    r._matchType = _matchType
    r.Set("match_type", _matchType)
    return nil
}

// MatchType Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchType() int64 {
    return r._matchType
}
// Num Setter
// 参赛号
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetNum(_num string) error {
    r._num = _num
    r.Set("num", _num)
    return nil
}

// Num Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetNum() string {
    return r._num
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAliuid() string {
    return r._aliuid
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispSign() string {
    return r._alispSign
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispTime() string {
    return r._alispTime
}
// MatchTime Setter
// 比赛日期 格式：Y-m-d
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchTime(_matchTime string) error {
    r._matchTime = _matchTime
    r.Set("match_time", _matchTime)
    return nil
}

// MatchTime Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchTime() string {
    return r._matchTime
}
// GunshotScore Setter
// 枪声成绩
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGunshotScore(_gunshotScore string) error {
    r._gunshotScore = _gunshotScore
    r.Set("gunshot_score", _gunshotScore)
    return nil
}

// GunshotScore Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGunshotScore() string {
    return r._gunshotScore
}
// GunshotRanking Setter
// 枪声排名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGunshotRanking(_gunshotRanking int64) error {
    r._gunshotRanking = _gunshotRanking
    r.Set("gunshot_ranking", _gunshotRanking)
    return nil
}

// GunshotRanking Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGunshotRanking() int64 {
    return r._gunshotRanking
}
// Speed Setter
// 平均配速
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSpeed(_speed string) error {
    r._speed = _speed
    r.Set("speed", _speed)
    return nil
}

// Speed Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSpeed() string {
    return r._speed
}
// Project Setter
// 展示用，比如：男子半程马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetProject(_project string) error {
    r._project = _project
    r.Set("project", _project)
    return nil
}

// Project Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetProject() string {
    return r._project
}
// SubScore Setter
// 比如马拉松 5KM 10KM 15KM分段成绩，json key->value 格式
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSubScore(_subScore string) error {
    r._subScore = _subScore
    r.Set("sub_score", _subScore)
    return nil
}

// SubScore Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSubScore() string {
    return r._subScore
}
// SubPoint Setter
// 比如马拉松 5KM 10KM 15KM分段时间点，json key->value 格式
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSubPoint(_subPoint string) error {
    r._subPoint = _subPoint
    r.Set("sub_point", _subPoint)
    return nil
}

// SubPoint Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSubPoint() string {
    return r._subPoint
}
