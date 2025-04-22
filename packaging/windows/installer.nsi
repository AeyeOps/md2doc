; NSIS installer script for md2docx
!define VERSION "${VERSION}"
OutFile "md2docx-${VERSION}-installer.exe"
InstallDir "$PROGRAMFILES64\\md2docx"
Page directory
Page instfiles

Section "Install md2docx"
    SetOutPath "$INSTDIR"
    File "md2docx_${VERSION}_windows_amd64.exe"
    CreateShortCut "$DESKTOP\\md2docx.lnk" "$INSTDIR\\md2docx.exe"
SectionEnd

Section "Uninstall"
    Delete "$INSTDIR\\md2docx.exe"
    Delete "$DESKTOP\\md2docx.lnk"
    RMDir "$INSTDIR"
SectionEnd