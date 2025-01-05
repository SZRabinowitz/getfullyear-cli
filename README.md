# getfullyear-cli
A cli client for getfullyear.com

## Features
- Versatile: This is good for many use cases. Use it when you forget the year, need to add a copyright footer to your cli, or you want to see trending ADs.
- Extensible: With many output options, including JSON, you can use this with whatever you want
- Compliant: This getfullyear.com client is fully compliant with their terms of service, so you can sleep well at night knowing you will not be sued. 


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
