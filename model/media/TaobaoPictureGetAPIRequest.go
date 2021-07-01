package media

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPictureGetAPIRequest
获取图片信息 API请求
taobao.picture.get

获取图片信息 */
type TaobaoPictureGetAPIRequest struct {
	model.Params
	// 图片ID
	_pictureId int64
	// 图片分类ID
	_pictureCategoryId int64
	// 是否删除,unfroze代表没有删除
	_deleted string
	// 图片标题,最大长度50字符,中英文都算一字符
	_title string
	// 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc
	_orderBy string
	// 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss
	_startDate string
	// 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss
	_endDate string
	// 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1
	_pageNo int64
	// 每页条数.每页返回最多返回100条,默认值40
	_pageSize int64
	// 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。
	_modifiedTime string
	// 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部
	_clientType string
	// 图片url查询接口
	_urls string
	// 是否获取https的链接
	_isHttps bool
}

// New
