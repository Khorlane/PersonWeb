cd /Projects/PersonWeb
@cls
@echo off
echo.
echo ****************************
echo Must be RUN AS ADMINISTRATOR
echo ****************************
@echo on
copy "C:\Projects\PersonWeb\models\*.*" "C:\Program Files\Go\src\models\*.*"
@if %ERRORLEVEL% == 0 goto :BuildIt
@echo.
@echo Copy FAILED.  Exited with status: %errorlevel%
@echo.
@pause
@goto :ExitScript
:BuildIt
del personweb.exe
go build
:ExitScript
@echo.
@echo Done
@echo.
@pause