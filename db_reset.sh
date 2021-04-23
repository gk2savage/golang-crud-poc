echo -e "\e[1;32mrunning db_reset.sh\e[0m"

spinner() {
    local i sp n
    sp='. . '
    n=${#sp}
    printf ' '
    while sleep 0.5; do
        printf  "${sp:i++%n:1}"
    done
}

printf 'setting up DB'
spinner &

sleep 3  # sleeping for 10 seconds is important work

kill "$!" # kill the spinner
printf '\n'


rm /home/gk/go/src/golang-crud-poc/database/qh.db
cp /home/gk/Downloads/backup/golang-crud-poc/database/qh.db /home/gk/go/src/golang-crud-poc/database/
echo "==============================DB Reset Done================================"
