import React, { useEffect, useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom';

export default function TwoFactorAuth() {
    const location = useLocation();
    const navigate = useNavigate();
    const state = location.state;
    const [completeAuthState,setCompleteAuthState] = useState<{location_id:number,user_id:number}>({
        location_id:0,
        user_id:0
    })
    
    useEffect(()=>{
        if ( state && state.user_id && state.location_id){
                setCompleteAuthState({
                    location_id: state.location_id,
                    user_id: state.user_id,
                })
        } else {
            alert('Invalid state')
            navigate('/');
        }

    },[])


    const onCompleteAuthClicked = (e:React.MouseEvent<HTMLButtonElement>)=>{
            e.preventDefault();
            fetch('http://localhost:8080/verify-auth',{
                method:"POST",
                body:JSON.stringify(completeAuthState)
            }).then(res =>res.json()).then(_=>{
                navigate('/home');
            }).catch(e=>{
                console.error('error in completing auth',e)
                alert('error in completing auth')
                navigate('/');
            })
    }
  return (
    <div className='text-center flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8'>
    <div className="sm:mx-auto sm:w-full sm:max-w-sm">
    <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900 w-50">
       Two Factor Auth
        </h2>
        <button
              onClick={onCompleteAuthClicked}
              type="submit"
              className="mt-10 flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
            Complete
            </button>
        </div>
    </div>
  )
}
