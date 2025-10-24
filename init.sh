# Shell Integration Code for Rocket

function rocket {
  # Decide how to handle commands based on the first argument
  case "$1" in
  # ---
  # Commands that need shell integration
  # ---
  new)
    # Run rocket-bin, capture its stdout, and 'eval' it
    eval "$(rocket-bin "$@")"
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
