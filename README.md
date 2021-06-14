# translate

This is a translation tool for translating content from input file into target language.

Currently support json, yaml, html and plain text formats.

The supported json format is:

```
{
    "key1": "value1",
    "key2": "value2",
    ...
}
```

The supported yaml format is:

```
key1: value1
key2: value2
...
```

## Installation

```
go get -u github.com/hyperjiang/translate
```

## Use Aliyun machine translation

Setup environment variables:

```
export ALI_REGION_ID="cn-hangzhou"
export ALI_ACCESS_KEY_ID="your-key"
export ALI_ACCESS_SECRET="your-secret"
```

```
translate aliyun -i "input-file" -o "output-file" -s source-language -t target-language
```
