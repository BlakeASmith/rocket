# Shell Integration Code for Rocket

function rocket {
  # Decide how to handle commands based on the first argument
  case "$1" in
  # ---
  # Commands that need shell integration
  # ---
  new)
    # Check if --no-go flag is present
    if [[ "$*" == *"--no-go"* ]]; then
      # Run rocket-bin directly without eval (no directory change)
      rocket-bin "$@"
    else
      # Run rocket-bin, capture its stdout, and 'eval' it (includes cd command)
      eval "$(rocket-bin "$@")"
    fi
    ;;
  goto)
    # Run rocket-bin, capture its stdout, and 'cd' to it
    cd "$(rocket-bin "$@")"
    ;;

  # ---
  # Commands that don't need shell integration
  # ---
  *)
    # Just execute rocket-bin directly.
    # Its stdout/stderr will pass through.
    rocket-bin "$@"
    ;;
  esac
}
