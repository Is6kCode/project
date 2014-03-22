/*
   @2014-03-21 将字段保存成Tab间隔的数据列
*/
package file

import(
	"os"
)

//公司头信息
type CompanyInfoItem struct {
	Order            string //序号
	CompanyName      string //单位名称
	ComRating        string //综合评级
	Deadline         string //有效日期
	ComQualification string //企业资质
	Level            string //级别
	Performance      string //业绩
	BidingDate       string //中标日期
	BidingPrice      string //中标价格
	ContractPrice    string //合同价
	CompleteDate     string //竣工日期
	Qualification    string //资质
}

//CompanyInfoItem 组织为Tab风格的串
func (infoItem CompanyInfoItem) FormatTabString() string {
	retStr := infoItem.Order + "\t" +
		infoItem.CompanyName + "\t" +
		infoItem.ComRating + "\t" +
		infoItem.Deadline + "\t" +
		infoItem.ComQualification + "\t" +
		infoItem.Level + "\t" +
		infoItem.Performance + "\t" +
		infoItem.BidingDate + "\t" +
		infoItem.BidingPrice + "\t" +
		infoItem.ContractPrice + "\t" +
		infoItem.CompleteDate + "\t" +
		infoItem.Qualification + "\t"
	return retStr
}

//将公司信息保存Tab分割的文件
func WriteToFile(FileName string, data []CompanyInfoItem) (err error) {
	//以读写方式打开或创建方式打开文件，且清空文件
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE,0644)
	if err != nil {
		return err
	}

	lineText := "序号\t单位名称\t综合评价分数\t有效期\t企业资质\t业绩名\t中标日期\t中标价(万元)\t竣工验收日期\t业绩对应的资质"
	file.WriteString(lineText)
	for _, infoItem := range data {
		lineText = "\r\n" + infoItem.FormatTabString()
		file.WriteString(lineText)
	}
	defer file.Close()
	return nil
}
