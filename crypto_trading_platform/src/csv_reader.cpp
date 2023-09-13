#include "csv_reader.hpp"
#include <exception>
#include <iostream>
#include <fstream>

CSVReader::CSVReader() {}

std::vector<OrderBookEntry> CSVReader::ReadCSV(std::string csv_file_path) {
    std::vector<OrderBookEntry> entries;
    std::string line;

    std::ifstream csv_file{csv_file_path};
    if (csv_file.is_open()) {
        std::cout << "File is opend \n";
    }
    else {
        std::cout << "Could not open file. Exiting!! \n";
        exit(-1);
    }

    while (std::getline(csv_file, line)) {
        try {
            OrderBookEntry obe = StringsToOBEObj(Tokenise(line, ','));
            entries.push_back(obe);
        } catch (const std::exception& e) {
            std::cout << "Bad value! " << line << std::endl;
            break;
        }
    }

    csv_file.close();

    return entries;
}

std::vector<std::string> CSVReader::Tokenise(std::string csv_line, char separator, char esc_seq) {
    std::vector<std::string> tokens;
    std::string token;
    auto start = csv_line.find_first_not_of(separator, 0);
    auto end = csv_line.find_first_of(separator, start);
    auto len = csv_line.length();

    if (esc_seq == ' ') {
        do {
            end = csv_line.find_first_of(separator, start);

            if (start == end) {
                if (start >= len) break;
                else {
                    start = end + 1;
                    continue;
                }
            }

            if (end != std::string::npos) token = csv_line.substr(start, end - start);
            else token = csv_line.substr(start, len - start);

            tokens.push_back(token);

            start = end + 1;
        } while (end != std::string::npos);
    }
    else {
        std::cout << "Code is in progress\n";
    }

    return tokens;
}

OrderBookEntry CSVReader::StringsToOBEObj(std::vector<std::string> tokens) {
    double price;
    double amount;

    // we have 5 tokens
    if (tokens.size() != 5) {
        std::cout << "Bad line" << std::endl;
        throw std::exception{};
    }

    try {
        // problem: stod reads till it encounters a non-numerical character
        // exception is not thrown unless
        price = std::stod(tokens[3]);
        amount = std::stod(tokens[4]);

    } catch (const std::exception& e) {
        std::cout << "Bad input! " << tokens[3] << std::endl;
        std::cout << "Bad input! " << tokens[4] << std::endl;
        throw e;
    }

    OrderBookEntry obe_object{price, amount, tokens[0], tokens[1], OrderBookEntry::StringToOBT(tokens[2])};

    return obe_object;
}
