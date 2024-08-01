@echo off
:loop
start /b "" "gep.exe" --tool 3
timeout /t 10 /nobreak >nul
goto loop
