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
    alispAppKey   string
    // 成绩(比赛用时)
    score   string
    // 组别 1001半程马拉松  1002全程马拉松
    matchGroup   int64
    // 国家
    country   string
    // 出生日期 格式：Y-m-d
    birthday   string
    // 手机号
    mobile   string
    // 证件ID
    cardId   string
    // 证件类型 1身份证 2军官证 4护照 8台胞证 16港澳通行证 32未设置  64 其他
    cardType   int64
    // 类型：1专业 2业余
    type   int64
    // 性别 0未知 1男 2女
    gender   int64
    // 姓名
    name   string
    // 排名
    ranking   int64
    // 比赛名（展示用）
    match   string
    // 比赛类型 1马拉松
    matchType   int64
    // 参赛号
    num   string
    // 阿里体育用户id
    aliuid   string
    // 接口签名
    alispSign   string
    // 时间戳精确到秒
    alispTime   string
    // 比赛日期 格式：Y-m-d
    matchTime   string
    // 枪声成绩
    gunshotScore   string
    // 枪声排名
    gunshotRanking   int64
    // 平均配速
    speed   string
    // 展示用，比如：男子半程马拉松
    project   string
    // 比如马拉松 5KM 10KM 15KM分段成绩，json key->value 格式
    subScore   string
    // 比如马拉松 5KM 10KM 15KM分段时间点，json key->value 格式
    subPoint   string
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
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// Score Setter
// 成绩(比赛用时)
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetScore(score string) error {
    r.score = score
    r.Set("score", score)
    return nil
}

// Score Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetScore() string {
    return r.score
}
// MatchGroup Setter
// 组别 1001半程马拉松  1002全程马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchGroup(matchGroup int64) error {
    r.matchGroup = matchGroup
    r.Set("match_group", matchGroup)
    return nil
}

// MatchGroup Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchGroup() int64 {
    return r.matchGroup
}
// Country Setter
// 国家
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

// Country Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCountry() string {
    return r.country
}
// Birthday Setter
// 出生日期 格式：Y-m-d
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

// Birthday Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetBirthday() string {
    return r.birthday
}
// Mobile Setter
// 手机号
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMobile() string {
    return r.mobile
}
// CardId Setter
// 证件ID
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCardId(cardId string) error {
    r.cardId = cardId
    r.Set("card_id", cardId)
    return nil
}

// CardId Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCardId() string {
    return r.cardId
}
// CardType Setter
// 证件类型 1身份证 2军官证 4护照 8台胞证 16港澳通行证 32未设置  64 其他
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCardType(cardType int64) error {
    r.cardType = cardType
    r.Set("card_type", cardType)
    return nil
}

// CardType Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCardType() int64 {
    return r.cardType
}
// Type Setter
// 类型：1专业 2业余
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetType() int64 {
    return r.type
}
// Gender Setter
// 性别 0未知 1男 2女
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGender(gender int64) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

// Gender Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGender() int64 {
    return r.gender
}
// Name Setter
// 姓名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetName() string {
    return r.name
}
// Ranking Setter
// 排名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetRanking(ranking int64) error {
    r.ranking = ranking
    r.Set("ranking", ranking)
    return nil
}

// Ranking Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetRanking() int64 {
    return r.ranking
}
// Match Setter
// 比赛名（展示用）
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatch(match string) error {
    r.match = match
    r.Set("match", match)
    return nil
}

// Match Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatch() string {
    return r.match
}
// MatchType Setter
// 比赛类型 1马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchType(matchType int64) error {
    r.matchType = matchType
    r.Set("match_type", matchType)
    return nil
}

// MatchType Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchType() int64 {
    return r.matchType
}
// Num Setter
// 参赛号
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetNum(num string) error {
    r.num = num
    r.Set("num", num)
    return nil
}

// Num Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetNum() string {
    return r.num
}
// Aliuid Setter
// 阿里体育用户id
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAliuid() string {
    return r.aliuid
}
// AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispSign() string {
    return r.alispSign
}
// AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispTime() string {
    return r.alispTime
}
// MatchTime Setter
// 比赛日期 格式：Y-m-d
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchTime(matchTime string) error {
    r.matchTime = matchTime
    r.Set("match_time", matchTime)
    return nil
}

// MatchTime Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchTime() string {
    return r.matchTime
}
// GunshotScore Setter
// 枪声成绩
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGunshotScore(gunshotScore string) error {
    r.gunshotScore = gunshotScore
    r.Set("gunshot_score", gunshotScore)
    return nil
}

// GunshotScore Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGunshotScore() string {
    return r.gunshotScore
}
// GunshotRanking Setter
// 枪声排名
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGunshotRanking(gunshotRanking int64) error {
    r.gunshotRanking = gunshotRanking
    r.Set("gunshot_ranking", gunshotRanking)
    return nil
}

// GunshotRanking Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGunshotRanking() int64 {
    return r.gunshotRanking
}
// Speed Setter
// 平均配速
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSpeed(speed string) error {
    r.speed = speed
    r.Set("speed", speed)
    return nil
}

// Speed Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSpeed() string {
    return r.speed
}
// Project Setter
// 展示用，比如：男子半程马拉松
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetProject(project string) error {
    r.project = project
    r.Set("project", project)
    return nil
}

// Project Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetProject() string {
    return r.project
}
// SubScore Setter
// 比如马拉松 5KM 10KM 15KM分段成绩，json key->value 格式
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSubScore(subScore string) error {
    r.subScore = subScore
    r.Set("sub_score", subScore)
    return nil
}

// SubScore Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSubScore() string {
    return r.subScore
}
// SubPoint Setter
// 比如马拉松 5KM 10KM 15KM分段时间点，json key->value 格式
func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSubPoint(subPoint string) error {
    r.subPoint = subPoint
    r.Set("sub_point", subPoint)
    return nil
}

// SubPoint Getter
func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSubPoint() string {
    return r.subPoint
}
