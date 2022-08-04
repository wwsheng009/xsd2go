# XSD2Go - Automatically generate golang xml parser based on XSD
![Build CI](https://github.com/GoComply/xsd2go/workflows/Build%20CI/badge.svg)
[![Lint CI](https://github.com/GoComply/xsd2go/actions/workflows/lint.yaml/badge.svg)](https://github.com/GoComply/xsd2go/actions/workflows/lint.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gocomply/xsd2go)](https://goreportcard.com/report/github.com/gocomply/xsd2go)
[![Gitter](https://badges.gitter.im/GoComply/community.svg)](https://gitter.im/GoComply/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gocomply/xsd2go)](https://pkg.go.dev/github.com/gocomply/xsd2go)

:warning: **You should run xsd2go, before ever importing `encoding/xml` to your project.** :warning:

You may want to start reading [blog introduction](http://isimluk.com/posts/2020/05/xsd2go-automatically-generate-golang-xml-parsers/) to this project.

## Motivation

Did you ever got to implement XML parser? Perhaps for atom, or scap? You may have got XSD
(XML Schema Definition) files to verify adherence to given xml application? Wouldn't it be
cool to automatically generate XML parser based on XSD definition? That's exactly what we
are up to here.

### Related projects:
 - ![Metaschema](https://github.com/gocomply/metaschema) - generate golang code based on NIST metaschema input
 - ![SCAP](https://github.com/gocomply/scap) - parsers of NIST SCAP family of standards

## Exemplary Usage

```
# Acquire latest some XSD file you want to convert - for instance XCCDF 1.2
git clone --depth 1 https://github.com/openscap/openscap
# Parse XSD schema and generate golang structs
./gocomply_xsd2go convert \
    --xmlns-override=http://cpe.mitre.org/language/2.0=cpe_language \
    .scap_schemas/schemas/xccdf/1.2/xccdf_1.2.xsd \
    github.com/gocomply/scap pkg/scap/models
```

## Installation

```
go get -u -v github.com/gocomply/xsd2go/cli/gocomply_xsd2go
```


## 原项目的一个缺点是无法支持在sequence上设置列表类型。比如下面的描述无法正确转换。
```
<xs:element name="GBSeq_feature-table" minOccurs="0">
        <xs:complexType>
          <xs:sequence minOccurs="0" maxOccurs="unbounded">
            <xs:element ref="GBFeature"/>
          </xs:sequence>
        </xs:complexType>
      </xs:element>
```


## 重新生成types.tmpl
修改types.tmpl，删除生成代码形成的空行
```
go install github.com/markbates/pkger/cmd/pkger@latest

//这样对types.tmpl修改才会生效
pkger -o /pkg/template

```