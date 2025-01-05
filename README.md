# getfullyear-cli
A cli client for getfullyear.com

## Features
- Versatile: This is good for many use cases. Use it when you forget the year, need to add a copyright footer to your cli, or you want to see trending ADs.
- Extensible: With many output options, including JSON, you can use this with whatever you want
- Compliant: This getfullyear.com client is fully compliant with their terms of service, so you can sleep well at night knowing you will not be sued.
- Cross Platform: Works on Windows, Linux, and MacOS. Want WASM? Look elsewhere!


Download Binaries in releases

## Usage:
```
C:\Users\Admin\Downloads\gofc>fullyear-cli
Red Bull: drink wing juice go zoom

--- Year Data ---
Year: 2024
Year String: 2024

C:\Users\Admin\Downloads\gofc>fullyear-cli --json
Hims: wellness simplified
{
  "year": 2025,
  "year_string": "2025",
  "sponsored_by": "Hims: wellness simplified"
}

C:\Users\Admin\Downloads\gofc>fullyear-cli --copyright "Script Kiddiez 4Life"
WhatsApp: blue ticks anxiety

--- Year Data ---
Year: 2025
Year String: 2025

Copyright © 2025 Script Kiddiez 4Life

C:\Users\Admin\Downloads\gofc>fullyear-cli --copyright "Software you can't trust" --json
Magnum: size matters
{
  "year": 2024,
  "year_string": "2024",
  "sponsored_by": "Magnum: size matters"
}

Copyright © 2024 Software you can't trust
```

## Contributing
- You can build this with the build.bat script
- We don't support Linux of Mac contributors
- If you are not on Windows, you can try Wine
- We take no responsibility if you get drunk

We have a few contribution rules:
- This project does NOT use git. You can only use the Github WebUI
- You can not use VIM to write code for this tool. We only allow Emacs.
- We are considering allowing Windows notepad, but we are still trying to load this file in the new Windows notepad to see how it looks. Once it opens we will provide an update.
- DO NOT WRITE TESTS: Due to constant bugs in the API, tests are prone to fail often. We cannot afford to be constantly tweaking our test cases.

## Known Issues
- getfullyear.com does not currently allow enterprise subscriptions because their serverless servers cannot handle the influx of new customers. Therfore, we cannot properly test this against enterprise mode. 
