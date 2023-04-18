# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [plugin/v1/plugin.proto](#plugin_v1_plugin-proto)
    - [PluginConfig](#plugin-v1-PluginConfig)
    - [PluginConfig.ConfigEntry](#plugin-v1-PluginConfig-ConfigEntry)
    - [PluginConfig.RequiresEntry](#plugin-v1-PluginConfig-RequiresEntry)
    - [PluginID](#plugin-v1-PluginID)
  
    - [HookName](#plugin-v1-HookName)
  
    - [GatewayDPluginService](#plugin-v1-GatewayDPluginService)
  
- [Scalar Value Types](#scalar-value-types)



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
| GetPluginConfig | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | GetPluginConfig returns the plugin config upon registration |
| OnConfigLoaded | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnConfigLoaded is called when the config is loaded from any config provider |
| OnNewLogger | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnNewLogger is called when a new logger is created |
| OnNewPool | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnNewPool is called when a new pool is created |
| OnNewClient | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnNewClient is called when a new client is created |
| OnNewProxy | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnNewProxy is called when a new proxy is created |
| OnNewServer | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnNewServer is called when a new server is created |
| OnSignal | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnSignal is called when an OS signal is received |
| OnRun | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnRun is called when the gatewayd server is running |
| OnBooting | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnBooting is called when the gatewayd server is booting |
| OnBooted | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnBooted is called when the gatewayd server is booted |
| OnOpening | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnOpening is called when the gatewayd server is opening |
| OnOpened | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnOpened is called when the gatewayd server is opened |
| OnClosing | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnClosing is called when the gatewayd server is closing |
| OnClosed | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnClosed is called when the gatewayd server is closed |
| OnTraffic | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnTraffic is called when the gatewayd server received traffic from a client |
| OnTrafficFromClient | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnTrafficFromClient is called when the gatewayd server is receiving traffic from a client |
| OnTrafficToServer | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnTrafficToServer is called when the gatewayd server is sending traffic to a server |
| OnTrafficFromServer | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnTrafficFromServer is called when the gatewayd server is receiving traffic from a server |
| OnTrafficToClient | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnTrafficToClient is called when the gatewayd server is sending traffic to a client |
| OnShutdown | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnShutdown is called when the gatewayd server is shutting down |
| OnTick | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnTick is called when the gatewayd server is ticking |
| OnHook | [.google.protobuf.Struct](#google-protobuf-Struct) | [.google.protobuf.Struct](#google-protobuf-Struct) | OnHook is called when the gatewayd server is calling a custom hook |

 



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

