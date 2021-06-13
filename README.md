# translate

Command line tool for translation.

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
