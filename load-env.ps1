# Load environment variables from .env file
function Load-EnvFile {
    param(
        [string]$Path = ".env"
    )
    
    if (Test-Path $Path) {
        Get-Content $Path | ForEach-Object {
            if ($_ -match "^\s*([^#][^=]*?)\s*=\s*(.*?)\s*$") {
                $name = $matches[1]
                $value = $matches[2]
                
                # Remove quotes if present
                $value = $value -replace '^["'']|["'']$'
                
                # Set environment variable
                [Environment]::SetEnvironmentVariable($name, $value, "Process")
                Write-Host "Set $name = $value" -ForegroundColor Green
            }
        }
        Write-Host "Environment variables loaded from $Path" -ForegroundColor Cyan
    } else {
        Write-Host "File $Path not found" -ForegroundColor Red
    }
}

# Usage: Load-EnvFile
# or: Load-EnvFile -Path "path/to/your/.env"
