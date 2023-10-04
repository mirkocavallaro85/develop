#include <iostream>
#include <string>
#include <curl/curl.h>
#include <nlohmann/json.hpp>

using json = nlohmann::json;

// Funzione di callback per scrivere i dati ricevuti in una variabile std::string
size_t WriteCallback(void *contents, size_t size, size_t nmemb, std::string *output) {
    size_t total_size = size * nmemb;
    output->append(static_cast<char*>(contents), total_size);
    return total_size;
}

int main() {
    // Inizializza libcurl
    CURL *curl = curl_easy_init();
    if (!curl) {
        std::cerr << "Errore inizializzando libcurl." << std::endl;
        return 1;
    }

    // Imposta l'URL della richiesta
    std::string url = "http://localhost:8080/example"; // Sostituisci con l'URL effettivo dell'API
    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());

    // Imposta la funzione di callback per scrivere i dati nella variabile response_data
    std::string response_data;
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, WriteCallback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response_data);

    // Effettua la richiesta HTTP
    CURLcode res = curl_easy_perform(curl);
    if (res != CURLE_OK) {
        std::cerr << "Errore durante la richiesta HTTP: " << curl_easy_strerror(res) << std::endl;
        return 1;
    }

    // Parsing del JSON
    try {
        json json_data = json::parse(response_data);
        std::cout << "Json: " << json_data << std::endl;
        // Ora puoi accedere ai dati JSON come desiderato
        std::string nome = json_data["nome"];

        // Esempio di output dei dati
        std::cout << "Nome: " << nome << std::endl;
    } catch (const std::exception& e) {
        std::cerr << "Errore durante il parsing del JSON: " << e.what() << std::endl;
        return 1;
    }

    // Rilascio delle risorse di libcurl
    curl_easy_cleanup(curl);

    return 0;
}