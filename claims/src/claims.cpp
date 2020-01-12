#include <eosio/eosio.hpp>
#include <eosio/system.hpp>

using namespace eosio;
using std::string;


class [[eosio::contract("claims")]]  claims : public contract {

public:
    using contract::contract;

    claims(name receiver, name code, datastream<const char *> ds) : contract(receiver, code, ds) {}

    /*
     * event entity
     * events for debug info, similar ETH events
     */
    struct [[eosio::table]] event {
        uint64_t    id;    //uniq id
        uint64_t    time;  //event time
        uint8_t     eId;   //event id
        string      log;   //log

        uint64_t primary_key() const { return id; }
    };

    struct [[eosio::table]] claimevent {
        uint64_t        id;         //uniq id
        uint64_t        createdAt;  //event time
        uint128_t       vinInt;     //vin indexed
        string          vin;        //vin str
        int32_t         gpsLt;     //lt without point
        int32_t         gpsLa;
        uint64_t primary_key() const { return id; }
        uint128_t get_secondary_1() const { return vinInt; }
    };

    struct [[eosio::table]] claim {
        uint64_t        id;             //uniq id
        uint64_t        clEventId;      //claim event
        uint8_t         formatVersion;  //data format version
        uint16_t        version;        //claim version
        uint64_t        createdAt;
        name            createdBy;      //police account
        std::vector<uint8_t>	data;           //serialized binary data for claim, can be changed depends on version field
        uint64_t primary_key() const { return id; }
        uint64_t get_secondary_1() const { return clEventId; }
    };

    //use for debug only
    [[eosio::action]]
    void hi(name user) {
        emit(1, "hi " + user.to_string());
    }

    [[eosio::action]]
    void addevent(uint64_t createdAt, int32_t gpsLt, int32_t gpsLa, uint128_t vinInt, string vin) {
        addClaimEvent(createdAt, vinInt, vin, gpsLt, gpsLa);
        //emit(1, vin);
    }

    //only for contract owner (self)
    [[eosio::action]]
    void delevents(name sender) {
        require_auth(sender);
        deleteAllClaimEvents();
    }

    //only for contract owner (self)
    [[eosio::action]]
    void delevents(name sender) {
        require_auth(sender);
        deleteAllEvents();
    }

private:
    typedef eosio::multi_index<"event"_n, event> eventTblType;
    typedef eosio::multi_index<"claimev"_n, claimevent, indexed_by<"byvin"_n, const_mem_fun<claimevent, uint128_t, &claimevent::get_secondary_1>>>  claimEventTblType;
    typedef eosio::multi_index<"claim"_n, claim, indexed_by<"byevent"_n, const_mem_fun<claim, uint64_t, &claim::get_secondary_1>>>  claimTblType;

    void emit(uint8_t eId, string log) {
        eventTblType eventTbl(get_self(), get_first_receiver().value);

        eventTbl.emplace(_self, [&](auto &event) {
            //increment id
            event.id = eventTbl.available_primary_key();
            event.time = eosio::current_block_time().block_timestamp_epoch;
            event.eId = eId;
            event.log = log;
        });
    }

    void addClaimEvent(uint64_t createdAt, uint128_t vinInt, string vin, int32_t gpsLt, int32_t gpsLa) {
        claimEventTblType claimeETbl(get_self(), get_first_receiver().value);

        claimeETbl.emplace(_self, [&](auto &evt) {
            //increment id
            evt.id = claimeETbl.available_primary_key();
            evt.createdAt = createdAt;
            evt.vinInt = vinInt;
            evt.vin = vin;
            evt.gpsLt = gpsLt;
            evt.gpsLa = gpsLa;
        });
    }

    void deleteAllClaimEvents() {
        claimEventTblType claimeETbl(get_self(), get_first_receiver().value);
        auto it = claimeETbl.begin();
        while (it != claimeETbl.end()) {
            it = claimeETbl.erase(it);
        }
    }

    void deleteAllEvents() {
        eventTblType eventTbl(get_self(), get_first_receiver().value);
        auto it = eventTbl.begin();
        while (it != eventTbl.end()) {
            it = eventTbl.erase(it);
        }
    }
};