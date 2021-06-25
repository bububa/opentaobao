package shop

// OpenApiSearchRequest 
type OpenApiSearchRequest struct {

    // 渠道版本 alipay,koubei,eleme,gaode,taobao
    App   string `json:"app,omitempty"`

    // 所选城市
    CityId   string `json:"city_id,omitempty"`

    // 客户端
    ClientOs   string `json:"client_os,omitempty"`

    // 客户端版本
    ClientVersion   string `json:"client_version,omitempty"`

    // 客户端上下文
    Context   string `json:"context,omitempty"`

    // 所选城市
    CurrentCity   string `json:"current_city,omitempty"`

    // 所选区
    CurrentDistrict   string `json:"current_district,omitempty"`

    // 省
    CurrentProvince   string `json:"current_province,omitempty"`

    // 固定值
    Forward   bool `json:"forward,omitempty"`

    // 纬度
    Latitude   string `json:"latitude,omitempty"`

    // 所在商圈
    LbsBusiAreaId   string `json:"lbs_busi_area_id,omitempty"`

    // 所在城市
    LbsCityId   string `json:"lbs_city_id,omitempty"`

    // 所在区
    LbsDistrictId   string `json:"lbs_district_id,omitempty"`

    // 经纬度精度
    LocationAccuracy   string `json:"location_accuracy,omitempty"`

    // 经度
    Longitude   string `json:"longitude,omitempty"`

    // 个数
    OriginalSize   int64 `json:"original_size,omitempty"`

    // OS 版本
    OsVersion   string `json:"os_version,omitempty"`

    // 扩展参数
    ParamsMap   string `json:"params_map,omitempty"`

    // 埋点ID
    PlaceId   string `json:"place_id,omitempty"`

    // 关键词
    Query   string `json:"query,omitempty"`

    // 请求ID
    RequestId   string `json:"request_id,omitempty"`

    // 场景吗
    SceneId   string `json:"scene_id,omitempty"`

    // 搜索ID
    SearchId   string `json:"search_id,omitempty"`

    // 菜单项
    SelectedMenus   string `json:"selected_menus,omitempty"`

    // sessionid
    SessionId   string `json:"session_id,omitempty"`

    // 个数
    Size   int64 `json:"size,omitempty"`

    // 排序
    Sort   string `json:"sort,omitempty"`

    // 来源埋点
    SrcSpm   string `json:"src_spm,omitempty"`

    // 分页
    Start   int64 `json:"start,omitempty"`

    // 搜索TOKEN
    TokenId   string `json:"token_id,omitempty"`

    // traceID
    Trace   string `json:"trace,omitempty"`

    // 用户ID
    Uid   string `json:"uid,omitempty"`

}
