# translate

[![](https://goreportcard.com/badge/github.com/hyperjiang/translate)](https://goreportcard.com/report/github.com/hyperjiang/translate)
[![License](https://img.shields.io/github/license/hyperjiang/translate.svg)](https://github.com/hyperjiang/translate)

This is a tool for translation and file format converting.

For file format, we currently support `json`, `yaml`, `properties` and `typescript`.

For translation client, we currently use Aliyun machine translation.

## Installation

```
go get -u github.com/hyperjiang/translate
```

## File converter

```
translate file -i "input-file" -o "output-file"
```

## Translator

For translation, we only support simple single-layer key-value format, e.g.

```
{
    "key1": "value1",
    "key2": "value2",
}
```

or

```
key1: value1
key2: value2
```

or

```
key1=value1
key2=value2
```

### Use Aliyun machine translation

Setup environment variables:

```
export ALI_REGION_ID="cn-hangzhou"
export ALI_ACCESS_KEY_ID="your-key"
export ALI_ACCESS_SECRET="your-secret"
```

```
translate aliyun -i "input-file" -o "output-file" -s source-language -t target-language
```
