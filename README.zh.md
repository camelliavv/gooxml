[English](./README.md) | **中文**

---

**gooxml** 是一个用于创建 Office Open XML 文档（.docx、.xlsx 和 .pptx）的库。它的目标是成为创建和编辑 docx/xlsx/pptx 文件的最兼容、性能最高的 Go 库。

需要 **go1.8+**，已在 1.8、1.9 和 tip 版本上测试通过。

[![License: AGPL v3](https://img.shields.io/badge/License-Dual%20AGPL%20v3/Commercial-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![GoDoc](https://godoc.org/baliance.com/gooxml?status.svg)](https://godoc.org/baliance.com/gooxml)
[![go 1.8+](https://img.shields.io/badge/go-1.8%2B-blue.svg)](http://golang.org)

![https://baliance.com/gooxml/](./_examples/preview.png "gooxml")

## 状态 ##

- 文档（docx）[Word]
	- 读取/写入/编辑
	- 格式化
	- 图片
	- 表格
- 电子表格（xlsx）[Excel]
 	- 读取/写入/编辑
 	- 单元格格式化，包括条件格式化
	- 单元格验证（下拉组合框、规则等）
    - 以 Excel 格式化后的形式获取单元格值（例如以 Excel 中显示的格式获取日期或数字）
 	- 公式计算（目前支持 100+ 函数，将根据需要添加更多）
 	- 嵌入图片
 	- 所有图表类型
- 演示文稿（pptx）[PowerPoint]
	- 从模板创建
	- 文本框/形状


## 性能 ##

最近有很多人关注电子表格创建/读取的性能数据，以下是 gooxml 在这个[基准测试](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/lots-of-rows)中的数据，该测试创建一个包含 3 万行、每行 100 列的表格。

    创建 30000 行 * 100 个单元格耗时 3.92506863 秒
    保存耗时 89 纳秒
    读取耗时 9.522383048 秒

创建速度相当快，由于不使用反射，保存非常迅速，读取稍慢一些。缺点是二进制文件较大（33MB），因为它包含了 DOCX/XLSX/PPTX 的所有生成结构体、序列化和反序列化代码。

## 安装 ##

    go get github.com/camelliavv/gooxml/
    go build -i github.com/camelliavv/gooxml/...

## 文档示例 ##

- [简单文本格式化](https://github.com/camelliavv/gooxml/tree/master/_examples/document/simple) 文本字体颜色、大小、高亮等。
- [自动生成目录](https://github.com/camelliavv/gooxml/tree/master/_examples/document/toc) 根据标题创建自动生成目录的文档标题
- [浮动图片](https://github.com/camelliavv/gooxml/tree/master/_examples/document/image) 在页面上绝对定位放置图片，支持不同的文本环绕方式。
- [页眉和页脚](https://github.com/camelliavv/gooxml/tree/master/_examples/document/header-footer) 创建包含页码的页眉和页脚。
- [多种页眉和页脚](https://github.com/camelliavv/gooxml/tree/master/_examples/document/header-footer-multiple) 根据文档部分使用不同的页眉和页脚。
- [行内表格](https://github.com/camelliavv/gooxml/tree/master/_examples/document/tables) 添加带边框和不带边框的表格。
- [使用现有 Word 文档作为模板](https://github.com/camelliavv/gooxml/tree/master/_examples/document/use-template) 打开文档作为模板，复用文档中创建的样式。
- [填写表单字段](https://github.com/camelliavv/gooxml/tree/master/_examples/document/fill-out-form) 打开包含嵌入式表单字段的文档，填写字段并将结果保存为新的填写表单。
- [编辑现有文档](https://github.com/camelliavv/gooxml/tree/master/_examples/document/edit-document) 打开现有文档并替换/删除文本，而不修改格式。

## 电子表格示例 ##
- [简单](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/simple) 一个包含几个单元格的简单表格
- [命名单元格](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/named-cells) 引用行和单元格的不同方式
- [单元格数字/日期/时间格式](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/number-date-time-formats) 创建具有各种数字/日期/时间格式的单元格
- [折线图](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/line-chart)/[3D 折线图](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/line-chart-3d) 折线图
- [柱状图](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/bar-chart) 柱状图
- [多个图表](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/multiple-charts) 单个表格上的多个图表
- [命名单元格范围](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/named-ranges) 命名单元格范围
- [合并单元格](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/merged) 合并和取消合并单元格
- [条件格式化](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/conditional-formatting) 单元格条件格式化、样式、渐变、图标、数据条
- [复杂](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/complex) 多个图表、自动筛选和条件格式化
- [边框](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/borders) 单个单元格边框和单元格范围周围的矩形边框。
- [验证](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/validation) 数据验证，包括组合框下拉菜单。
- [冻结行/列](https://github.com/camelliavv/gooxml/tree/master/_examples/spreadsheet/freeze-rows-cols) 带有冻结标题列和行的表格

## 演示文稿示例 ##

- [简单文本框](https://github.com/camelliavv/gooxml/tree/master/_examples/presentation/simple) 简单的文本框和形状
- [图片](https://github.com/camelliavv/gooxml/tree/master/_examples/presentation/image) 简单的图片插入
- [模板](https://github.com/camelliavv/gooxml/tree/master/_examples/presentation/use-template/simple) 从模板创建演示文稿

## 原始类型 ##

OOXML 规范非常庞大，为整个规范创建友好的 API 是一项非常耗时的任务。这个库试图在提供易于使用的 API 用于创建 OOXML 文档的常见用例的同时，允许用户在库的 API 未涵盖特定用例时回退到原始文档操作。

原始基于 XML 的类型位于 ```schema/``` 目录中。这些类型可以通过包装类型的 ```X()``` 方法访问，该方法返回原始类型。

例如，库目前没有用于设置文档背景颜色的 API。但是，通过编辑文档的 ```CT_Background``` 元素可以很容易地手动完成。

    doc := document.New()
    doc.X().Background = wordprocessingml.NewCT_Background()
	doc.X().Background.ColorAttr = &wordprocessingml.ST_HexColor{}
	doc.X().Background.ColorAttr.ST_HexColorRGB = color.RGB(50, 50, 50).AsRGBString()

### 贡献指南 ###

[![CLA assistant](https://cla-assistant.io/readme/badge/baliance/gooxml)](https://cla-assistant.io/baliance/gooxml)

所有贡献者在代码被审查和合并之前必须签署贡献者许可协议。


### 许可 ###

这个库采用双重许可提供。根据 AGPLv3 条款，它可以免费使用。如果您想在闭源项目中使用这个库，请联系 sales@unidoc.io。

开源版本和商业版本在功能上没有区别。鼓励您使用开源版本在购买商业许可之前评估这个库。

