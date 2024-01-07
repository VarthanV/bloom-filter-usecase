import React from 'react'

export default function TwoFactorAuth() {
    
  return (
    <div className='text-center flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8'>
    <div className="sm:mx-auto sm:w-full sm:max-w-sm">
    <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900 w-50">
       Two Factor Auth
        </h2>
        <button
              onClick={(e) => {
                e.preventDefault();
              
              }}
              type="submit"
              className="mt-10 flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            >
            Complete
            </button>
        </div>
    </div>
  )
}
