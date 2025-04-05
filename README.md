# Scraping Portals

A Go application for scraping news articles from local news portals using Go and MySQL.

## Features

- Scrapes articles from multiple news portals:
  - Telegrafi
  - GazetaExpress
  - IndeksOnline
- Categorizes articles automatically
- Stores data in MySQL database
- Supports image handling
- API integration for data posting

## Prerequisites

- Go 1.14 or higher
- MySQL database
- Required environment variables set up for API and database access

## Environment Variables

The following environment variables need to be configured:

### API Credentials
- `api_post_url`: API endpoint for posting articles
- `api_image_url`: API endpoint for image uploads
- `api_telegrafi_post_user`: Telegrafi API username
- `api_telegrafi_post_password`: Telegrafi API password
- `api_gazetaexpress_post_user`: GazetaExpress API username
- `api_gazetaexpress_post_password`: GazetaExpress API password
- `api_indeksonline_post_user`: IndeksOnline API username
- `api_indeksonline_post_password`: IndeksOnline API password

### Database Configuration
- `mysql_portals_username`: MySQL username
- `mysql_portals_password`: MySQL password
- `mysql_portals_host`: MySQL host address
- `mysql_portals_schema`: MySQL database schema name

## Usage

Run the scraper for a specific portal using the `-portal` flag:

```bash
go run main.go -portal telegrafi
go run main.go -portal gazetaexpress
go run main.go -portal indeksonline
```

## Categories

The scraper supports the following article categories:
- Lajme (1)
- Sport (2)
- Magazina/Showbiz (3)
- Teknologji (4)
- Kuriozitete (5)
- Shendetesi (6)
- Ekonomi (7)
