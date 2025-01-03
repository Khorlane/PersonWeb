cd /Projects/PersonWeb
@cls
@echo off
echo.
@echo on
del personweb.exe
@if %ERRORLEVEL% == 0 goto :BuildIt
@echo.
@echo Delete FAILED.  Exited with status: %errorlevel%
@echo.
@pause
@goto :ExitScript
:BuildIt
go build
:ExitScript
@echo.
@echo Done
@echo.
@pause