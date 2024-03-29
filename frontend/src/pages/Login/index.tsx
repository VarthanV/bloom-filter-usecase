import React, { useEffect, useState } from "react";
import Dropdown from "../../components/Dropdown";
import { DropdownOption } from "../../types";
import { useNavigate } from "react-router-dom";

export default function Login() {
  const [users, setUsers] = useState<DropdownOption[] | null>(null);
  const [locations, setLocations] = useState<DropdownOption[] | null>(null);
  const [state, setSelectedOption] = useState<{
    user_id: number;
    location_id: number;
  }>({
    user_id: 0,
    location_id: 0,
  });
  const [isAuthenticating, setIsAuthenticating] = useState(false);
  const navigate = useNavigate();



  const login = () => {
    setIsAuthenticating(true);
  };

  const displaySelectedLocation = () => {
    if (state.location_id && locations) {
      const location = locations.filter(
        (item) => item.value === state.location_id
      );
      if (location) {
        return location[0].option;
      }
      return "";
    }
  };

  const displaySelectedUser = () => {
    if (state.user_id && users) {
      const user = users.filter(
        (item) => item.value === state.user_id
      );
      if (user) {
        return user[0].option;
      }
      return "";
    }
  };

  const getLocations = () => {
    fetch("http://localhost:8080/locations")
      .then((res) => res.json())
      .then((data) => {
        setLocations(data);
      })
      .catch((e) => {
        console.error("error in getting locations ", e);
        alert("error in getting locations");
      });
  };

  const getUsers = () => {
    fetch("http://localhost:8080/users")
      .then((res) => res.json())
      .then((data) => {
        setUsers(data);
      })
      .catch((e) => {
        console.error("error in getting users ", e);
        alert("error in getting users");
      });
  };

  useEffect(() => {
    getLocations();
    getUsers();
  }, []);

  useEffect(() => {
    if (isAuthenticating) {
      fetch("http://localhost:8080/authenticate", {
        method: "POST",
        body: JSON.stringify(state),
      })
        .then((res) => res.json())
        .then((data) => {
         switch(data.goto){
            case 'homepage':
                navigate('/home');
                break
            case 'two_factor_auth':
                navigate('/challenge',{state:state})
                break
         }
        })
        .catch((e) => {
          console.error("error in getting users ", e);
          alert("error in getting users");
        });
    }
  }, [isAuthenticating]);

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-sm">
        <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
          Bloom filter Login Simulator
        </h2>

        {!isAuthenticating ? (
          <div>
            {locations && users ? (
              <div className="text-center flex content-center items-center mt-5">
                {/* Location */}
                <Dropdown
                  title="Location"
                  className="mr-3"
                  options={locations}
                  onOptionClick={(val) => {
                    setSelectedOption({
                      ...state,
                      location_id: val,
                    });
                  }}
                />

                {/* Users */}
                <Dropdown
                  title="Users"
                  options={users}
                  onOptionClick={(val) => {
                    setSelectedOption({
                      ...state,
                      user_id: val,
                    });
                  }}
                />
              </div>
            ) : (
              <p className="text-md text-center mt-3">Loading...</p>
            )}

            <div className="text-center mt-5">
              <p className="text-lg">
                <b>Location Selected : </b> {displaySelectedLocation()}{" "}
              </p>
              <p className="text-lg">
                <b> User Selected : </b> {displaySelectedUser()}{" "}
              </p>
            </div>
          </div>
        ) : (
          <p className="text-center text-lg font-bold"> Loading...</p>
        )}

        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <div>
            <button
              onClick={(e) => {
                e.preventDefault();
                login();
              }}
              type="submit"
              className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
              Sign in
            </button>
          </div>
        </div>

    
      </div>
    </div>
  );
}
