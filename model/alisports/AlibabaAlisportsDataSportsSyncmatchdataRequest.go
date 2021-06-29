package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育数据中心用户比赛数据同步接口 APIRequest
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

func NewAlibabaAlisportsDataSportsSyncmatchdataRequest() *AlibabaAlisportsDataSportsSyncmatchdataRequest{
    return &AlibabaAlisportsDataSportsSyncmatchdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetApiMethodName() string {
    return "alibaba.alisports.data.sports.syncmatchdata"
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetScore(score string) error {
    r.score = score
    r.Set("score", score)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetScore() string {
    return r.score
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchGroup(matchGroup int64) error {
    r.matchGroup = matchGroup
    r.Set("match_group", matchGroup)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchGroup() int64 {
    return r.matchGroup
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCountry() string {
    return r.country
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetBirthday() string {
    return r.birthday
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMobile() string {
    return r.mobile
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCardId(cardId string) error {
    r.cardId = cardId
    r.Set("card_id", cardId)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCardId() string {
    return r.cardId
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetCardType(cardType int64) error {
    r.cardType = cardType
    r.Set("card_type", cardType)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetCardType() int64 {
    return r.cardType
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetType() int64 {
    return r.type
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGender(gender int64) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGender() int64 {
    return r.gender
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetName() string {
    return r.name
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetRanking(ranking int64) error {
    r.ranking = ranking
    r.Set("ranking", ranking)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetRanking() int64 {
    return r.ranking
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatch(match string) error {
    r.match = match
    r.Set("match", match)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatch() string {
    return r.match
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchType(matchType int64) error {
    r.matchType = matchType
    r.Set("match_type", matchType)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchType() int64 {
    return r.matchType
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetNum(num string) error {
    r.num = num
    r.Set("num", num)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetNum() string {
    return r.num
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAliuid() string {
    return r.aliuid
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetMatchTime(matchTime string) error {
    r.matchTime = matchTime
    r.Set("match_time", matchTime)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetMatchTime() string {
    return r.matchTime
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGunshotScore(gunshotScore string) error {
    r.gunshotScore = gunshotScore
    r.Set("gunshot_score", gunshotScore)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGunshotScore() string {
    return r.gunshotScore
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetGunshotRanking(gunshotRanking int64) error {
    r.gunshotRanking = gunshotRanking
    r.Set("gunshot_ranking", gunshotRanking)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetGunshotRanking() int64 {
    return r.gunshotRanking
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSpeed(speed string) error {
    r.speed = speed
    r.Set("speed", speed)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSpeed() string {
    return r.speed
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetProject(project string) error {
    r.project = project
    r.Set("project", project)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetProject() string {
    return r.project
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSubScore(subScore string) error {
    r.subScore = subScore
    r.Set("sub_score", subScore)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSubScore() string {
    return r.subScore
}

func (r *AlibabaAlisportsDataSportsSyncmatchdataRequest) SetSubPoint(subPoint string) error {
    r.subPoint = subPoint
    r.Set("sub_point", subPoint)
    return nil
}

func (r AlibabaAlisportsDataSportsSyncmatchdataRequest) GetSubPoint() string {
    return r.subPoint
}

