# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [plugin/v1/struct.proto](#plugin_v1_struct-proto)
    - [ListValue](#plugin-v1-ListValue)
    - [Struct](#plugin-v1-Struct)
    - [Struct.FieldsEntry](#plugin-v1-Struct-FieldsEntry)
    - [Value](#plugin-v1-Value)
  
    - [NullValue](#plugin-v1-NullValue)
  
- [plugin/v1/plugin.proto](#plugin_v1_plugin-proto)
    - [PluginConfig](#plugin-v1-PluginConfig)
    - [PluginConfig.ConfigEntry](#plugin-v1-PluginConfig-ConfigEntry)
    - [PluginConfig.RequiresEntry](#plugin-v1-PluginConfig-RequiresEntry)
    - [PluginID](#plugin-v1-PluginID)
  
    - [HookName](#plugin-v1-HookName)
  
    - [GatewayDPluginService](#plugin-v1-GatewayDPluginService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="plugin_v1_struct-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## plugin/v1/struct.proto



<a name="plugin-v1-ListValue"></a>

### ListValue
`ListValue` is a wrapper around a repeated field of values.

The JSON representation for `ListValue` is JSON array.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [Value](#plugin-v1-Value) | repeated | Repeated field of dynamically typed values. |






<a name="plugin-v1-Struct"></a>

### Struct
`Struct` represents a structured data value, consisting of fields
which map to dynamically typed values. In some languages, `Struct`
might be supported by a native representation. For example, in
scripting languages like JS a struct is represented as an
object. The details of that representation are described together
with the proto support for the language.

The JSON representation for `Struct` is JSON object plus the
bytes type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [Struct.FieldsEntry](#plugin-v1-Struct-FieldsEntry) | repeated | Unordered map of dynamically typed values. |






<a name="plugin-v1-Struct-FieldsEntry"></a>

### Struct.FieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Value](#plugin-v1-Value) |  |  |






<a name="plugin-v1-Value"></a>

### Value
`Value` represents a dynamically typed value which can be either
null, a number, a string, a boolean, a recursive struct value, or a
list of values. A producer of value is expected to set one of these
variants. Absence of any variant indicates an error.

The JSON representation for `Value` is JSON value.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| null_value | [NullValue](#plugin-v1-NullValue) |  | Represents a null value. |
| number_value | [double](#double) |  | Represents a double value. |
| string_value | [string](#string) |  | Represents a string value. |
| bool_value | [bool](#bool) |  | Represents a boolean value. |
| struct_value | [Struct](#plugin-v1-Struct) |  | Represents a structured value. |
| list_value | [ListValue](#plugin-v1-ListValue) |  | Represents a repeated `Value`. |
| bytes_value | [bytes](#bytes) |  | Represents a bytes value. |





 


<a name="plugin-v1-NullValue"></a>

### NullValue
`NullValue` is a singleton enumeration to represent the null value for the
`Value` type union.

The JSON representation for `NullValue` is JSON `null`.

| Name | Number | Description |
| ---- | ------ | ----------- |
| NULL_VALUE | 0 | Null value. |


 

 

 



<a name="plugin_v1_plugin-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## plugin/v1/plugin.proto



<a name="plugin-v1-PluginConfig"></a>

### PluginConfig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [PluginID](#plugin-v1-PluginID) |  |  |
| description | [string](#string) |  |  |
| authors | [string](#string) | repeated |  |
| license | [string](#string) |  |  |
| project_url | [string](#string) |  |  |
| config | [PluginConfig.ConfigEntry](#plugin-v1-PluginConfig-ConfigEntry) | repeated | internal and external config options |
| hooks | [HookName](#plugin-v1-HookName) | repeated | hooks it attaches to |
| requires | [PluginConfig.RequiresEntry](#plugin-v1-PluginConfig-RequiresEntry) | repeated | required plugins |
| tags | [string](#string) | repeated |  |
| categories | [string](#string) | repeated |  |






<a name="plugin-v1-PluginConfig-ConfigEntry"></a>

### PluginConfig.ConfigEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="plugin-v1-PluginConfig-RequiresEntry"></a>

### PluginConfig.RequiresEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="plugin-v1-PluginID"></a>

### PluginID



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| version | [string](#string) |  |  |
| remote_url | [string](#string) |  |  |
| checksum | [string](#string) |  |  |





 


<a name="plugin-v1-HookName"></a>

### HookName


| Name | Number | Description |
| ---- | ------ | ----------- |
| HOOK_NAME_UNSPECIFIED | 0 |  |
| HOOK_NAME_ON_CONFIG_LOADED | 1 |  |
| HOOK_NAME_ON_NEW_LOGGER | 2 |  |
| HOOK_NAME_ON_NEW_POOL | 3 |  |
| HOOK_NAME_ON_NEW_CLIENT | 4 |  |
| HOOK_NAME_ON_NEW_PROXY | 5 |  |
| HOOK_NAME_ON_NEW_SERVER | 6 |  |
| HOOK_NAME_ON_SIGNAL | 7 |  |
| HOOK_NAME_ON_RUN | 8 |  |
| HOOK_NAME_ON_BOOTING | 9 |  |
| HOOK_NAME_ON_BOOTED | 10 |  |
| HOOK_NAME_ON_OPENING | 11 |  |
| HOOK_NAME_ON_OPENED | 12 |  |
| HOOK_NAME_ON_CLOSING | 13 |  |
| HOOK_NAME_ON_CLOSED | 14 |  |
| HOOK_NAME_ON_TRAFFIC | 15 |  |
| HOOK_NAME_ON_TRAFFIC_FROM_CLIENT | 16 |  |
| HOOK_NAME_ON_TRAFFIC_TO_SERVER | 17 |  |
| HOOK_NAME_ON_TRAFFIC_FROM_SERVER | 18 |  |
| HOOK_NAME_ON_TRAFFIC_TO_CLIENT | 19 |  |
| HOOK_NAME_ON_SHUTDOWN | 20 |  |
| HOOK_NAME_ON_TICK | 21 |  |
| HOOK_NAME_ON_HOOK | 22 |  |


 

 


<a name="plugin-v1-GatewayDPluginService"></a>

### GatewayDPluginService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPluginConfig | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | GetPluginConfig returns the plugin config upon registration |
| OnConfigLoaded | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnConfigLoaded is called when the config is loaded from any config provider |
| OnNewLogger | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnNewLogger is called when a new logger is created |
| OnNewPool | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnNewPool is called when a new pool is created |
| OnNewClient | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnNewClient is called when a new client is created |
| OnNewProxy | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnNewProxy is called when a new proxy is created |
| OnNewServer | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnNewServer is called when a new server is created |
| OnSignal | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnSignal is called when an OS signal is received |
| OnRun | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnRun is called when the gatewayd server is running |
| OnBooting | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnBooting is called when the gatewayd server is booting |
| OnBooted | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnBooted is called when the gatewayd server is booted |
| OnOpening | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnOpening is called when the gatewayd server is opening |
| OnOpened | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnOpened is called when the gatewayd server is opened |
| OnClosing | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnClosing is called when the gatewayd server is closing |
| OnClosed | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnClosed is called when the gatewayd server is closed |
| OnTraffic | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnTraffic is called when the gatewayd server received traffic from a client |
| OnTrafficFromClient | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnTrafficFromClient is called when the gatewayd server is receiving traffic from a client |
| OnTrafficToServer | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnTrafficToServer is called when the gatewayd server is sending traffic to a server |
| OnTrafficFromServer | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnTrafficFromServer is called when the gatewayd server is receiving traffic from a server |
| OnTrafficToClient | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnTrafficToClient is called when the gatewayd server is sending traffic to a client |
| OnShutdown | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnShutdown is called when the gatewayd server is shutting down |
| OnTick | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnTick is called when the gatewayd server is ticking |
| OnHook | [Struct](#plugin-v1-Struct) | [Struct](#plugin-v1-Struct) | OnHook is called when the gatewayd server is calling a custom hook |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

