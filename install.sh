#!/bin/bash

# Rocket CLI Installer
# This script builds and installs the rocket CLI tool with shell integration

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BINARY_NAME="rocket-bin"
INSTALL_DIR="/usr/local/bin"
SHELL_RC=""

# Detect shell and appropriate RC file
detect_shell() {
  case "$SHELL" in
  */zsh)
    SHELL_RC="$HOME/.zshrc"
    ;;
  */bash)
    # Check for bashrc or profile
    if [[ -f "$HOME/.bashrc" ]]; then
      SHELL_RC="$HOME/.bashrc"
    elif [[ -f "$HOME/.bash_profile" ]]; then
      SHELL_RC="$HOME/.bash_profile"
    else
      SHELL_RC="$HOME/.bashrc"
    fi
    ;;
  *)
    echo -e "${RED}Unsupported shell: $SHELL${NC}"
    echo -e "${RED}Please manually add the integration code to your shell profile.${NC}"
    exit 1
    ;;
  esac
}

# Check if running as root/sudo for system installation
check_permissions() {
  if [[ "$INSTALL_DIR" == "/usr/local/bin" ]] && [[ ! -w "$INSTALL_DIR" ]]; then
    if [[ "$EUID" -eq 0 ]]; then
      echo -e "${YELLOW}Running as root - installing to system location${NC}"
    else
      echo -e "${YELLOW}Need sudo access to install to $INSTALL_DIR${NC}"
      echo -e "${YELLOW}Installing to $HOME/.local/bin instead${NC}"
      INSTALL_DIR="$HOME/.local/bin"
      mkdir -p "$INSTALL_DIR"
      export PATH="$INSTALL_DIR:$PATH"
    fi
  fi
}

# Build the binary
build_binary() {
  echo -e "${BLUE}Building rocket CLI...${NC}"
  if ! command -v go &>/dev/null; then
    echo -e "${RED}Go is not installed. Please install Go first.${NC}"
    exit 1
  fi

  go build -o "$BINARY_NAME" .
  echo -e "${GREEN}Binary built successfully${NC}"
}

# Install binary
install_binary() {
  echo -e "${BLUE}Installing binary to $INSTALL_DIR...${NC}"

  if [[ "$INSTALL_DIR" == "/usr/local/bin" ]] && [[ "$EUID" -ne 0 ]]; then
    sudo cp "$BINARY_NAME" "$INSTALL_DIR/"
  else
    cp "$BINARY_NAME" "$INSTALL_DIR/"
  fi

  chmod +x "$INSTALL_DIR/$BINARY_NAME"
  echo -e "${GREEN}Binary installed to $INSTALL_DIR/$BINARY_NAME${NC}"
}

# Add shell integration
add_shell_integration() {
  # Check if already installed
  if grep -q "rocket-bin init" "$SHELL_RC" 2>/dev/null; then
    echo -e "${GREEN}Shell integration already configured in $SHELL_RC${NC}"
    return
  fi

  echo -e "${BLUE}Shell Integration Setup${NC}"
  echo ""
  echo -e "${YELLOW}To enable rocket commands in your shell, you need to add shell integration.${NC}"
  echo -e "${YELLOW}This will allow you to use 'rocket' commands directly.${NC}"
  echo ""
  echo -e "${GREEN}Option 1 - Temporary (current session only):${NC}"
  echo -e "  eval \"\$($INSTALL_DIR/$BINARY_NAME init)\""
  echo ""
  echo -e "${GREEN}Option 2 - Permanent (add to your shell profile):${NC}"
  echo -e "  echo 'eval \"\$($INSTALL_DIR/$BINARY_NAME init)\"' >> $SHELL_RC"
  echo -e "  source $SHELL_RC"
  echo ""

  # Ask for permission
  read -p "Would you like me to add the permanent integration to $SHELL_RC? (y/N): " -n 1 -r
  echo ""

  if [[ $REPLY =~ ^[Yy]$ ]]; then
    # Add eval command to shell RC
    echo "" >>"$SHELL_RC"
    echo "# Rocket CLI integration" >>"$SHELL_RC"
    echo "eval \"\$($INSTALL_DIR/$BINARY_NAME init)\"" >>"$SHELL_RC"
    echo -e "${GREEN}Shell integration added to $SHELL_RC${NC}"
    echo -e "${YELLOW}Please run 'source $SHELL_RC' to activate${NC}"
  else
    echo -e "${YELLOW}Shell integration not added. You can manually add it later.${NC}"
    echo -e "${YELLOW}For temporary use, run: eval \"\$($INSTALL_DIR/$BINARY_NAME init)\"${NC}"
  fi
}

# Main installation process
main() {
  echo -e "${BLUE}🚀 Rocket CLI Installer${NC}"
  echo ""

  detect_shell
  check_permissions
  build_binary
  install_binary
  add_shell_integration

  echo ""
  echo -e "${GREEN}Installation complete!${NC}"
  echo -e "${GREEN}Binary installed to $INSTALL_DIR/$BINARY_NAME${NC}"
}

# Run main function
main "$@"
