#!/usr/bin/env pwsh

# Quick setup script for DID API environment variables
Write-Host "Setting up environment variables for DID API..." -ForegroundColor Cyan

# App Configuration
$env:APP_NAME = "did-api"
$env:APP_PORT = "8080"
Write-Host "✓ App configuration set" -ForegroundColor Green

# Database Configuration
$env:DB_HOST = "localhost"
$env:DB_PORT = "5432"
$env:DB_USER = "postgres"  # Change this to your actual username
$env:DB_PASS = "password"  # Change this to your actual password
$env:DB_SCHEMA = "did_db"
$env:DB_MAX_IDLE_CONN = "10"
$env:DB_MAX_OPEN_CONN = "100"
$env:DB_MAX_CONN_LIFETIME = "3600"
$env:DB_DEBUG = "true"  # Set to true for development
Write-Host "✓ Database configuration set" -ForegroundColor Green

# Logger Configuration
$env:LOG_LEVEL = "debug"  # Use debug for development
$env:LOG_FILE_PATH = "logs/app.log"
Write-Host "✓ Logger configuration set" -ForegroundColor Green

Write-Host "Environment variables setup complete!" -ForegroundColor Cyan
Write-Host "You can now run your application with: go run main.go" -ForegroundColor Yellow

# Display current environment variables
Write-Host "`nCurrent environment variables:" -ForegroundColor Cyan
Write-Host "APP_NAME: $env:APP_NAME"
Write-Host "APP_PORT: $env:APP_PORT"
Write-Host "DB_HOST: $env:DB_HOST"
Write-Host "DB_PORT: $env:DB_PORT"
Write-Host "DB_USER: $env:DB_USER"
Write-Host "DB_SCHEMA: $env:DB_SCHEMA"
Write-Host "LOG_LEVEL: $env:LOG_LEVEL"
