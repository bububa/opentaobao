# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.2] - 2021-06-28

### Added

- Support self made metadata in patch which not exists in downloaded metadata

### Changed

### Deprecated

### Removed

### Fixed

- wrong decode error_response for xml format.

## [1.1.0] - 2021-06-28

### Added

- Support "HMac" sign method. use client.SetSignMethod(method model.SignMethod) to switch sign method.
- Support "XML" api format. use client.SetAPIFormat(format model.APIFormat) to switch api response format.
- Add Example (tbk.TaobaoTbkItemGet).
- Add "MapData", "Results", "RpcResult" into conflict_models.json. 

### Changed

- Changed model.tpl and response.tpl to support "XML" api format.
- For "JSON" api format will always use simplify=true.
- Optimized README.md display.
- APIResponse will not have "Response" field and add "RequestId" in each Response.

### Deprecated

### Removed

- Remove "RequestId" from CommonResponse.

### Fixed

- can't decode api response.

### Security

## [1.0.0] - 2021-06-25

This is the first release of opentaobao go library.
It contains api metadata downloader and generator. 
