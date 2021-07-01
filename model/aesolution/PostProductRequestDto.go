package aesolution

// PostProductRequestDto 结构体
type PostProductRequestDto struct {
	// aliexpress product Id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// group id, you can get group list from aliexpress.product.productgroups.get
	GroupId int64 `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// extra params. Configured some special products
	ExtraParams string `json:"extra_params,omitempty" xml:"extra_params,omitempty"`
	// marketing images for product. Currently supported 2 types: 1 represents 3:4 rectangle(resolution at least 750*1000) image while 2 represents 1:1 square image(Resolution at least 800*800). The image url needs to be obtained via uploading the image through Aliexpress API: aliexpress.photobank.redefining.uploadimageforsdk(https://developers.aliexpress.com/en/doc.htm?docId=30186&docType=2)
	MarketingImages []MarketImageDto `json:"marketing_images,omitempty" xml:"marketing_images>market_image_dto,omitempty"`
	// Aliexpress leaf category ID obtained through https://developers.aliexpress.com/en/doc.htm?docId=46042&docType=2
	AliexpressCategoryId int64 `json:"aliexpress_category_id,omitempty" xml:"aliexpress_category_id,omitempty"`
	// Deprecated. Please use aliexpress_category_id
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// merchant's brand name
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// Main image that represents the product. The url should be accesible and there is a meximum of 6 pictures. The url can point to a seller's server or to AliExpress photobank. In order to obtain more information about the photobank and how to upload images, please visit the following page: https://developers.aliexpress.com/en/doc.htm?docId=30186&docType=2
	MainImageUrlsList []string `json:"main_image_urls_list,omitempty" xml:"main_image_urls_list>string,omitempty"`
	// There are 4 types of how to fill in the content of each element in this attribute list. 1). When filling in the standard dropdown/multi-dropdown attributes, fill in "aliexpress_attribute_name_id" and "aliexpress_attribute_value_id"(For multi-dropdown, splitting them into multiple elements) 2). When filling in the attributes in which the value needs to be manually entered, fill in "aliexpress_attribute_name_id" and "attribute_value" in the element. 3). There exists a special kind of "aliexpress_attribute_value_id" of which the value represents for "Other". When encoutering this scenario, please fill in "aliexpress_attribute_name_id", "aliexpress_attribute_value_id" and "attribute_value". 4). Besides the three types mentioned above, if the seller would like to fully customized all the atttributes, fill in "attribute name" and "attribute_value" in the element.
	AttributeList []AttributeDto `json:"attribute_list,omitempty" xml:"attribute_list>attribute_dto,omitempty"`
	// If specified this field, all the previous skus will be replaced by the new skus.
	SkuInfoList []SkuInfoDto `json:"sku_info_list,omitempty" xml:"sku_info_list>sku_info_dto,omitempty"`
	// Weight of the product, including package. Measured in Kilograms (Kg) with a maximum 500 and minumum 0.001
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// indicate when the inventory of a specific product will be deducted. place_order_withhold: the inventory is deducted just after the order is placed by customer. payment_success_deduct: the stock is deducted after the payment is done successfully by the customer.
	InventoryDeductionStrategy string `json:"inventory_deduction_strategy,omitempty" xml:"inventory_deduction_strategy,omitempty"`
	// Package height measured in centimeters (cm). Maximum 700 cm, minumum: 0.01cm
	PackageHeight int64 `json:"package_height,omitempty" xml:"package_height,omitempty"`
	// Package length, measured in centimeters (cm). Maximum 700 cm, minumum: 0.01cm
	PackageLength int64 `json:"package_length,omitempty" xml:"package_length,omitempty"`
	// Package width measured in centimeters (cm). Maximum 700 cm, minumum: 0.01cm
	PackageWidth int64 `json:"package_width,omitempty" xml:"package_width,omitempty"`
	// multi country price configuration
	MultiCountryPriceConfiguration *MultiCountryPriceConfigurationDto `json:"multi_country_price_configuration,omitempty" xml:"multi_country_price_configuration,omitempty"`
	// freight template ID. After the merchant has created an freight template in the backend: freighttemplate.aliexpress.com, the id could be obtained in the backend directly or thourgh the API: aliexpress.freight.redefining.listfreighttemplate
	FreightTemplateId int64 `json:"freight_template_id,omitempty" xml:"freight_template_id,omitempty"`
	// refer to the preparation period of merchant before the package could be dispatched to the customer.
	ShippingLeadTime int64 `json:"shipping_lead_time,omitempty" xml:"shipping_lead_time,omitempty"`
	// service policy id, which could be set and obtained in the seller's background.
	ServicePolicyId int64 `json:"service_policy_id,omitempty" xml:"service_policy_id,omitempty"`
	// merchant's size chart id, more used in the category of shoes and clothes.
	SizeChartId int64 `json:"size_chart_id,omitempty" xml:"size_chart_id,omitempty"`
	// List for multi language subject. To learn how to set this field, please refer to the document:https://developers.aliexpress.com/en/doc.htm?docId=108976&docType=1
	MultiLanguageSubjectList []SingleLanguageTitleDto `json:"multi_language_subject_list,omitempty" xml:"multi_language_subject_list>single_language_title_dto,omitempty"`
	// List for multi language description. To learn how to set this field, please refer to the document:https://developers.aliexpress.com/en/doc.htm?docId=108976&docType=1
	MultiLanguageDescriptionList []SingleLanguageDescriptionDto `json:"multi_language_description_list,omitempty" xml:"multi_language_description_list>single_language_description_dto,omitempty"`
	// Deprecated. Please use  multi_language_subject_list. 1-128 length, support multi language. Check the field "language" to find the supported languages.
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// Deprecated, please use product description, support html format and multi languages. Check the field language to find the supported languages.
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// The main language of the product. Aliexpress will depend on the subject and detail in main language to translate to other languages, if not filled in by sellers. Main language could not be modified after product has been uploaded. Support: en(English) ru(Russian) es(Spanish) fr(French) it(Italian) tr(Turkish) pt(Portuguese) de(German) nl(Dutch) in(Indonesian) ar(Arabic) ja(Japanese) ko(Korean) th(Thai) vi(Vietnamese) iw(Hebrew)
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// Product Unit ID, Most common-used ID: 100000015 piece/pieces; 100000000:bag/bags; 100000001:barrel/barrels; 100000002:bushel/bushels; 100078580:carton; 100078581:centimeter; 100000003:cubic meter; 100000004:dozen; 100078584:feet; 100000005:gallon; 100000006:gram; 100078587:inch; 100000007:kilogram; 100078589:kiloliter; 100000008:kilometer; 100078559:liter/liters; 100000009:long ton; 100000010:meter; 100000011:metric ton; 100078560:milligram; 100078596:milliliter; 100078597:millimeter; 100000012:ounce; 100000014:pack/packs; 100000013:pair; 100000016:pound; 100078603:quart; 100000017:set/sets; 100000018:short ton; 100078606:square feet; 100078607:square inch; 100000019:square meter; 100078609:square yard; 100000020:ton; 100078558:yard/yards
	ProductUnit int64 `json:"product_unit,omitempty" xml:"product_unit,omitempty"`
}
