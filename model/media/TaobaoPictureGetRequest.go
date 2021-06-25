package media

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取图片信息 APIRequest
taobao.picture.get

获取图片信息
*/
type TaobaoPictureGetRequest struct {
    model.Params

    // 图片ID
    pictureId   int64 

    // 图片分类ID
    pictureCategoryId   int64 

    // 是否删除,unfroze代表没有删除
    deleted   string 

    // 图片标题,最大长度50字符,中英文都算一字符
    title   string 

    // 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
    orderBy   string 

    // 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
    startDate   string 

    // 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
    endDate   string 

    // 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
    pageNo   int64 

    // 每页条数.每页返回最多返回100条,默认值40
    pageSize   int64 

    // 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
    modifiedTime   string 

    // 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
    clientType   string 

    // 图片url查询接口
    urls   string 

    // 是否获取https的链接
    isHttps   bool 

}

func NewTaobaoPictureGetRequest() *TaobaoPictureGetRequest{
    return &TaobaoPictureGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPictureGetRequest) GetApiMethodName() string {
    return "taobao.picture.get"
}

func (r TaobaoPictureGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPictureGetRequest) SetPictureId(pictureId int64) error {
    r.pictureId = pictureId
    r.Set("picture_id", pictureId)
    return nil
}

func (r TaobaoPictureGetRequest) GetPictureId() int64 {
    return r.pictureId
}

func (r *TaobaoPictureGetRequest) SetPictureCategoryId(pictureCategoryId int64) error {
    r.pictureCategoryId = pictureCategoryId
    r.Set("picture_category_id", pictureCategoryId)
    return nil
}

func (r TaobaoPictureGetRequest) GetPictureCategoryId() int64 {
    return r.pictureCategoryId
}

func (r *TaobaoPictureGetRequest) SetDeleted(deleted string) error {
    r.deleted = deleted
    r.Set("deleted", deleted)
    return nil
}

func (r TaobaoPictureGetRequest) GetDeleted() string {
    return r.deleted
}

func (r *TaobaoPictureGetRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoPictureGetRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoPictureGetRequest) SetOrderBy(orderBy string) error {
    r.orderBy = orderBy
    r.Set("order_by", orderBy)
    return nil
}

func (r TaobaoPictureGetRequest) GetOrderBy() string {
    return r.orderBy
}

func (r *TaobaoPictureGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r TaobaoPictureGetRequest) GetStartDate() string {
    return r.startDate
}

func (r *TaobaoPictureGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r TaobaoPictureGetRequest) GetEndDate() string {
    return r.endDate
}

func (r *TaobaoPictureGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoPictureGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoPictureGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoPictureGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoPictureGetRequest) SetModifiedTime(modifiedTime string) error {
    r.modifiedTime = modifiedTime
    r.Set("modified_time", modifiedTime)
    return nil
}

func (r TaobaoPictureGetRequest) GetModifiedTime() string {
    return r.modifiedTime
}

func (r *TaobaoPictureGetRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

func (r TaobaoPictureGetRequest) GetClientType() string {
    return r.clientType
}

func (r *TaobaoPictureGetRequest) SetUrls(urls string) error {
    r.urls = urls
    r.Set("urls", urls)
    return nil
}

func (r TaobaoPictureGetRequest) GetUrls() string {
    return r.urls
}

func (r *TaobaoPictureGetRequest) SetIsHttps(isHttps bool) error {
    r.isHttps = isHttps
    r.Set("is_https", isHttps)
    return nil
}

func (r TaobaoPictureGetRequest) GetIsHttps() bool {
    return r.isHttps
}

