#pragma once

#include "orderbookentry.hpp"
#include <string>
#include <vector>

class CSVReader {
public:
    CSVReader();
    static std::vector<OrderBookEntry> ReadCSV(std::string csv_file_path);

private:
    static std::vector<std::string> Tokenise(std::string, char, char esc_seq = ' ');
    static OrderBookEntry StringsToOBEObj(std::vector<std::string> tokens);
};
