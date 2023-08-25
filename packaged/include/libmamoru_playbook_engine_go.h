/*! \file */
/*******************************************
 *                                         *
 *  File auto-generated by `::safer_ffi`.  *
 *                                         *
 *  Do not manually edit this file.        *
 *                                         *
 *******************************************/

#ifndef __RUST_PLAYBOOK_ENGINE_C__
#define __RUST_PLAYBOOK_ENGINE_C__
#ifdef __cplusplus
extern "C" {
#endif


#include <stdbool.h>

/** <No documentation available> */
typedef struct FfiJsonResult {
    /** <No documentation available> */
    bool is_error;

    /** <No documentation available> */
    char * error_message;

    /** <No documentation available> */
    char * result_json;
} FfiJsonResult_t;

/** \brief
 *  Frees a validation result
 */
void
ffi_drop_json_result (
    FfiJsonResult_t result);

/** \brief
 *  Resumes a playbook
 *  Arguments are passed as JSON strings
 *  Date is RFC3339
 */
FfiJsonResult_t
ffi_playbook_resume (
    char const * datetime,
    char const * run_id,
    char const * confirmations_json);

/** \brief
 *  Starts a playbook
 *  Arguments are passed as JSON strings
 *  Date is RFC3339
 */
FfiJsonResult_t
ffi_playbook_start (
    char const * datetime,
    char const * run_id,
    char const * playbook_json,
    char const * trigger_json);

/** \brief
 *  Validates a playbook
 */
FfiJsonResult_t
ffi_validate_playbook (
    char const * playbook_json);


#ifdef __cplusplus
} /* extern \"C\" */
#endif

#endif /* __RUST_PLAYBOOK_ENGINE_C__ */
