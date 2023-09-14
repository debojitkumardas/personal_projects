#pragma once

#include "orderbookentry.hpp"
#include <string>
#include <vector>

class Helper {
private:
    std::vector<OrderBookEntry> orders;

public:
    Helper();
    /** Call this to start the simulation by providing the source of data */
    void Init(std::string file_path);

private:
    void LoadOrderBook(std::string);
    void PrintMenu();
    void PrintHelp();
    void PrintMarketStats();
    void EnterOffer();
    void EnterBid();
    void PrintWallet();
    void GotoNextTimeframe();
    int GetUserOption();
    void ProcessUserOption(int user_option);
};
